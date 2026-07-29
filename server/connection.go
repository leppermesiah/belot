package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"

	"belot/engine"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // LAN/self-hosted use, no auth needed
}

// Keepalive timing. Without this, an idle connection (nothing sent for
// a while during bidding/thinking time) can get silently dropped by an
// intermediate proxy/NAT/router idle timeout - the ping/pong traffic
// keeps the connection visibly "alive" to anything watching for
// activity, and lets us detect a genuinely dead connection quickly
// instead of the browser eventually timing out on its own.
const (
	pongWait   = 30 * time.Second
	pingPeriod = (pongWait * 8) / 10 // must stay well under pongWait
	writeWait  = 10 * time.Second
)

type inMsg struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	TargetScore int    `json:"targetScore"`
	Action      string `json:"action"` // pass | suit | notrump | alltrump | contra | reconto
	Suit        string `json:"suit"`
	Card        string `json:"card"`
	Team        string `json:"team"`     // cherry | malina
	Kind        string `json:"kind"`     // announce kind, for declare messages
	HighRank    string `json:"highRank"` // for declare messages
	Value       int    `json:"value"`    // for declare messages
	Category    string `json:"category"` // "sequence" | "carre" | "belot", for declare messages
	TeamName    string `json:"teamName"`
}

func ServeWS(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	c := &Client{conn: conn, send: make(chan []byte, 32)}
	go c.writePump()
	c.readLoop(hub)
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()
	for {
		select {
		case msg, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) readLoop(hub *Hub) {
	defer func() {
		if c.room != nil {
			c.room.mu.Lock()
			c.room.removeClientLocked(c)
			c.room.broadcastRoomState()
			c.room.mu.Unlock()
		}
		c.conn.Close()
	}()

	for {
		_, raw, err := c.conn.ReadMessage()
		if err != nil {
			return
		}
		var m inMsg
		if err := json.Unmarshal(raw, &m); err != nil {
			continue
		}

		if c.room == nil {
			c.handleLobbyMsg(hub, m)
			continue
		}
		c.handleGameMsg(m)
	}
}

func (c *Client) handleLobbyMsg(hub *Hub, m inMsg) {
	switch m.Type {
	case "create_room":
		room := hub.Create(m.TargetScore)
		c.name = m.Name
		room.mu.Lock()
		room.addClientLocked(c)
		c.room = room
		room.broadcastRoomState()
		room.mu.Unlock()
	case "join_room":
		room := hub.Get(m.Code)
		if room == nil {
			b, _ := json.Marshal(map[string]interface{}{"type": "error", "message": "room_not_found"})
			c.send <- b
			return
		}
		if room.NumJoined >= 4 {
			b, _ := json.Marshal(map[string]interface{}{"type": "error", "message": "room_full"})
			c.send <- b
			return
		}
		c.name = m.Name
		room.mu.Lock()
		room.addClientLocked(c)
		c.room = room
		room.broadcastRoomState()
		if room.Started && room.Match != nil {
			// A match is already in progress (this is a reconnect after
			// a refresh/dropped connection, not a fresh lobby join) -
			// send everyone the current game state so the rejoining
			// player lands back on the table instead of being stuck on
			// the lobby screen with a Start button that no longer does
			// anything (the server already refuses to re-start a match
			// that's already running).
			room.broadcastGameState()
		}
		room.mu.Unlock()
	}
}

func (c *Client) handleGameMsg(m inMsg) {
	room := c.room
	room.mu.Lock()
	defer room.mu.Unlock()

	switch m.Type {
	case "choose_team":
		c.handleChooseTeam(m)

	case "set_team_name":
		c.handleSetTeamName(m)

	case "declare":
		if room.Match == nil || room.Match.Current == nil || m.Kind == "" {
			return
		}
		team := c.seat % 2
		switch m.Category {
		case "sequence":
			room.DeclaredSequence[team] = true
		case "carre":
			room.DeclaredCarre[team] = true
		case "belot":
			room.DeclaredBelot[team] = true
		}
		room.broadcastRaw(map[string]interface{}{
			"type":     "declared",
			"player":   c.seat,
			"kind":     m.Kind,
			"suit":     m.Suit,
			"highRank": m.HighRank,
			"value":    m.Value,
		})

	case "start_game":
		if room.Started || room.NumJoined < 4 {
			return
		}
		cherry, malina := 0, 0
		for _, t := range room.TeamChoice {
			switch t {
			case "cherry":
				cherry++
			case "malina":
				malina++
			}
		}
		if cherry != 2 || malina != 2 {
			room.broadcastError(c.seat, "teams_incomplete_before_start")
			return
		}

		// Reseat: cherry -> game seats 0 & 2 (TeamA), malina -> game
		// seats 1 & 3 (TeamB) - teammates must sit opposite each other
		// for the engine's seat-parity team logic to line up with the
		// team the players actually picked.
		oldClients := room.Clients
		oldTeamChoice := room.TeamChoice
		var newClients [4]*Client
		var newTeamChoice [4]string
		cherrySlot, malinaSlot := 0, 1
		for i, cl := range oldClients {
			if oldTeamChoice[i] == "cherry" {
				newClients[cherrySlot] = cl
				newTeamChoice[cherrySlot] = "cherry"
				cl.seat = cherrySlot
				cherrySlot += 2
			} else {
				newClients[malinaSlot] = cl
				newTeamChoice[malinaSlot] = "malina"
				cl.seat = malinaSlot
				malinaSlot += 2
			}
		}
		room.Clients = newClients
		// Keep TeamChoice indexed by the NEW seat positions - otherwise
		// a second start_game (a real rematch via "Нова игра", not a
		// fresh deal) would reseat against stale pre-reseat indices and
		// scramble the teams that were already correctly seated.
		room.TeamChoice = newTeamChoice

		var names [4]string
		for i, cl := range room.Clients {
			names[i] = cl.name
		}
		room.Match = engine.NewMatch(names, room.TargetScore)
		room.Match.StartHand()
		room.Started = true
		room.broadcastGameState()

	case "bid":
		c.handleBid(m)

	case "play_card":
		c.handlePlayCard(m)

	case "ready_for_next_hand":
		c.handleReadyForNextHand()
	}
}

// handleChooseTeam must be called with room.mu already held (it is,
// since the caller holds it for the whole handleGameMsg switch). This
// is what makes the "3 players can't grab the same team even if they
// click at the same instant" guarantee hold: every choose_team message
// is processed one at a time under this lock, so the second and third
// simultaneous clicks simply see the team already full.
func (c *Client) handleChooseTeam(m inMsg) {
	room := c.room
	if room.Started {
		return
	}
	if m.Team != "cherry" && m.Team != "malina" {
		room.broadcastError(c.seat, "invalid_team")
		return
	}
	if room.TeamChoice[c.seat] == m.Team {
		return // no-op, already on this team
	}
	count := 0
	for i, t := range room.TeamChoice {
		if i != c.seat && t == m.Team {
			count++
		}
	}
	if count >= 2 {
		room.broadcastError(c.seat, "team_full")
		return
	}
	room.TeamChoice[c.seat] = m.Team
	room.broadcastRoomState()
	room.maybePromptTeamNames()
}

// handleSetTeamName must be called with room.mu already held.
func (c *Client) handleSetTeamName(m inMsg) {
	room := c.room
	name := strings.TrimSpace(m.TeamName)
	if name == "" {
		return
	}
	if len(name) > 24 {
		name = name[:24]
	}
	teamStr := room.TeamChoice[c.seat]
	idx := -1
	if teamStr == "cherry" {
		idx = 0
	} else if teamStr == "malina" {
		idx = 1
	}
	if idx == -1 {
		return
	}
	room.TeamNames[idx] = name
	room.broadcastRoomState()
}

// handleBid must be called with room.mu already held.
func (c *Client) handleBid(m inMsg) {
	room := c.room
	match := room.Match
	if match == nil || match.Bidding == nil {
		return
	}
	b := match.Bidding
	var err error
	switch m.Action {
	case "pass":
		err = b.Pass(c.seat)
	case "suit":
		var s engine.Suit
		s, err = strToSuit(m.Suit)
		if err == nil {
			err = b.CallSuit(c.seat, s)
		}
	case "notrump":
		err = b.CallNoTrump(c.seat)
	case "alltrump":
		err = b.CallAllTrump(c.seat)
	case "contra":
		err = b.Contra(c.seat)
	case "reconto":
		err = b.Reconto(c.seat)
	default:
		err = errUnknownAction
	}
	if err != nil {
		room.broadcastError(c.seat, err.Error())
		return
	}

	if b.Done {
		redealt := match.AfterBiddingResolved()
		if redealt {
			room.broadcastGameState() // new hand dealt, bidding restarts
			return
		}
	}
	room.broadcastGameState()
}

// handlePlayCard must be called with room.mu already held.
func (c *Client) handlePlayCard(m inMsg) {
	room := c.room
	match := room.Match
	if match == nil || match.Current == nil {
		return
	}
	card, err := strToCard(m.Card)
	if err != nil {
		room.broadcastError(c.seat, err.Error())
		return
	}
	prevTricks := len(match.Current.TricksWon)
	if err := match.Current.PlayCard(c.seat, card); err != nil {
		room.broadcastError(c.seat, err.Error())
		return
	}
	if len(match.Current.TricksWon) > prevTricks {
		completed := match.Current.TricksWon[len(match.Current.TricksWon)-1]
		room.broadcastTrickComplete(completed, match.Current.TrickLeader)
	}

	if match.Current.Phase == engine.PhaseScoring {
		declared := engine.DeclaredAnnounces{
			Sequence: map[engine.TeamID]bool{engine.TeamA: room.DeclaredSequence[0], engine.TeamB: room.DeclaredSequence[1]},
			Carre:    map[engine.TeamID]bool{engine.TeamA: room.DeclaredCarre[0], engine.TeamB: room.DeclaredCarre[1]},
			Belot:    map[engine.TeamID]bool{engine.TeamA: room.DeclaredBelot[0], engine.TeamB: room.DeclaredBelot[1]},
		}
		res, over, winner := match.FinishHandAndCheckOverWithDeclared(declared)
		room.DeclaredSequence = [2]bool{}
		room.DeclaredCarre = [2]bool{}
		room.DeclaredBelot = [2]bool{}
		room.broadcastGameState()
		room.broadcastHandResult(res)

		if over {
			room.broadcastMatchOver(winner)
			room.Started = false
			return
		}
		room.ReadyNext = [4]bool{}
		return
	}
	room.broadcastGameState()
}

// handleReadyForNextHand must be called with room.mu already held. Once
// every seated player has confirmed, it starts the next hand - this is
// what makes everyone see the same full hand recap and click "Готово"
// before play resumes, instead of the game silently jumping ahead.
func (c *Client) handleReadyForNextHand() {
	room := c.room
	if room.Match == nil || !room.Started || room.Match.Current == nil {
		return
	}
	if room.Match.Current.Phase != engine.PhaseFinished {
		return
	}
	room.ReadyNext[c.seat] = true
	room.broadcastReadyStatus()

	allReady := true
	for i := 0; i < 4; i++ {
		if room.Clients[i] != nil && !room.ReadyNext[i] {
			allReady = false
			break
		}
	}
	if allReady {
		room.ReadyNext = [4]bool{}
		room.Match.StartHand()
		room.broadcastGameState()
	}
}

var errUnknownAction = &protoErr{"unknown bid action"}

type protoErr struct{ msg string }

func (e *protoErr) Error() string { return e.msg }
