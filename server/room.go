package server

import (
	"encoding/json"
	"log"
	"math/rand"
	"sync"

	"github.com/gorilla/websocket"

	"belot/engine"
)

type Client struct {
	conn *websocket.Conn
	seat int
	name string
	send chan []byte
	room *Room
}

type Room struct {
	Code             string
	TargetScore      int
	mu               sync.Mutex
	Clients          [4]*Client
	NumJoined        int
	TeamChoice       [4]string // "" | "cherry" | "malina", indexed by join slot (same index as Clients)
	Match            *engine.Match
	Started          bool
	ReadyNext        [4]bool   // who has clicked "Готово" to advance past the just-finished hand
	TeamNames        [2]string // ["Отбор А","Отбор Б"] by default; cherry=index0, malina=index1
	TeamNamePrompted [2]bool   // whether the naming prompt has been sent for that team
	TeamNamerSeat    [2]int    // which join-slot was randomly picked to name that team (-1 if none)
	DeclaredSequence [2]bool   // per TeamID: did they click "declare" for a sequence this hand
	DeclaredCarre    [2]bool   // per TeamID: did they click "declare" for a carre this hand
	DeclaredBelot    [2]bool   // per TeamID: did they declare belot this hand
}

// maybePromptTeamNames must be called with room.mu already held. Once
// a team fills to exactly 2 members, it randomly picks ONE of those
// two players and asks them to name the team - that name sticks for
// the rest of the match. Fires at most once per team (guarded by
// TeamNamePrompted), even if team membership later changes.
func (r *Room) maybePromptTeamNames() {
	var cherrySeats, malinaSeats []int
	for i, t := range r.TeamChoice {
		switch t {
		case "cherry":
			cherrySeats = append(cherrySeats, i)
		case "malina":
			malinaSeats = append(malinaSeats, i)
		}
	}
	if len(cherrySeats) == 2 && !r.TeamNamePrompted[0] {
		r.TeamNamePrompted[0] = true
		seat := cherrySeats[rand.Intn(2)]
		r.TeamNamerSeat[0] = seat
		r.sendTeamNamePrompt(seat, "cherry")
	}
	if len(malinaSeats) == 2 && !r.TeamNamePrompted[1] {
		r.TeamNamePrompted[1] = true
		seat := malinaSeats[rand.Intn(2)]
		r.TeamNamerSeat[1] = seat
		r.sendTeamNamePrompt(seat, "malina")
	}
}

func (r *Room) sendTeamNamePrompt(seat int, team string) {
	if r.Clients[seat] == nil {
		return
	}
	b, err := json.Marshal(map[string]interface{}{"type": "prompt_team_name", "team": team})
	if err != nil {
		return
	}
	select {
	case r.Clients[seat].send <- b:
	default:
	}
}

func newRoomCode() string {
	const chars = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func NewRoom(targetScore int) *Room {
	if targetScore <= 0 {
		targetScore = 151
	}
	return &Room{
		Code:          newRoomCode(),
		TargetScore:   targetScore,
		TeamNames:     [2]string{"Отбор А", "Отбор Б"},
		TeamNamerSeat: [2]int{-1, -1},
	}
}

// addClientLocked seats a client in the first free slot. Must be
// called with room.mu already held. Returns the seat index, or -1 if
// the room is full.
func (r *Room) addClientLocked(c *Client) int {
	for i := 0; i < 4; i++ {
		if r.Clients[i] == nil {
			r.Clients[i] = c
			c.seat = i
			r.NumJoined++
			return i
		}
	}
	return -1
}

// addClient is the locking entry point for callers that don't need to
// do anything else atomically alongside the seating.
func (r *Room) addClient(c *Client) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.addClientLocked(c)
}

// removeClientLocked must be called with room.mu already held.
func (r *Room) removeClientLocked(c *Client) {
	if r.Clients[c.seat] == c {
		r.Clients[c.seat] = nil
		r.NumJoined--
	}
}

func (r *Room) removeClient(c *Client) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.removeClientLocked(c)
}

func (r *Room) broadcastRaw(msg map[string]interface{}) {
	b, err := json.Marshal(msg)
	if err != nil {
		log.Println("marshal error:", err)
		return
	}
	for _, c := range r.Clients {
		if c != nil {
			select {
			case c.send <- b:
			default:
				log.Println("client send buffer full, dropping message for seat", c.seat)
			}
		}
	}
}

func (r *Room) broadcastRoomState() {
	names := make([]interface{}, 4)
	for i, c := range r.Clients {
		if c != nil {
			names[i] = map[string]interface{}{"seat": i, "name": c.name, "team": r.TeamChoice[i]}
		} else {
			names[i] = nil
		}
	}
	for _, c := range r.Clients {
		if c == nil {
			continue
		}
		msg := map[string]interface{}{
			"type":             "room_state",
			"code":             r.Code,
			"players":          names,
			"started":          r.Started,
			"yourSlot":         c.seat,
			"teamNames":        []string{r.TeamNames[0], r.TeamNames[1]},
			"teamNamePrompted": []bool{r.TeamNamePrompted[0], r.TeamNamePrompted[1]},
			"teamNamerSeat":    []int{r.TeamNamerSeat[0], r.TeamNamerSeat[1]},
		}
		b, err := json.Marshal(msg)
		if err != nil {
			log.Println("marshal error:", err)
			continue
		}
		select {
		case c.send <- b:
		default:
			log.Println("client send buffer full, dropping message for seat", c.seat)
		}
	}
}

// broadcastGameState sends a personalized state to each seated client -
// each player sees only their own hand; others' hands are shown as
// counts, per the "server is the source of truth" design.
func (r *Room) broadcastGameState() {
	m := r.Match
	if m == nil || m.Current == nil {
		return
	}
	g := m.Current

	handCounts := make([]int, 4)
	for i, p := range g.Players {
		handCounts[i] = len(p.Hand)
	}

	trick := make([]interface{}, 0, len(g.CurrentTrick))
	for _, pc := range g.CurrentTrick {
		trick = append(trick, map[string]interface{}{
			"player": pc.PlayerID,
			"card":   cardToStr(pc.Card),
		})
	}

	var contractView interface{}
	if g.Contract != nil {
		contractView = map[string]interface{}{
			"type":     contractTypeToStr(g.Contract.Type),
			"suit":     suitToStr(g.Contract.Suit),
			"callerId": g.Contract.CallerID,
			"contra":   g.Contract.Contra,
			"reconto":  g.Contract.Reconto,
		}
	}

	playerNames := make([]string, 4)
	for i, cl := range r.Clients {
		if cl != nil {
			playerNames[i] = cl.name
		}
	}

	for _, c := range r.Clients {
		if c == nil {
			continue
		}
		hand := make([]string, len(g.Players[c.seat].Hand))
		for i, card := range g.Players[c.seat].Hand {
			hand[i] = cardToStr(card)
		}

		var legalMoves []string
		if g.Phase == engine.PhasePlaying && g.Turn == c.seat {
			moves, err := g.LegalMoves(c.seat)
			if err == nil {
				legalMoves = make([]string, len(moves))
				for i, mv := range moves {
					legalMoves[i] = cardToStr(mv)
				}
			}
		}

		var myAnnounces []interface{}
		if g.Phase == engine.PhasePlaying && len(g.Players[c.seat].Hand) == 8 {
			for _, a := range engine.DetectAnnounces(g, c.seat, g.Players[c.seat].Hand) {
				category := "sequence"
				if a.Category == engine.CategoryCarre {
					category = "carre"
				}
				myAnnounces = append(myAnnounces, map[string]interface{}{
					"label":    describeAnnounce(a),
					"value":    a.Value,
					"category": category,
				})
			}
		}

		msg := map[string]interface{}{
			"type":         "game_state",
			"phase":        phaseToStr(g.Phase),
			"handNumber":   m.HandNumber,
			"dealer":       g.Dealer,
			"turn":         g.Turn,
			"trump":        suitToStr(g.Trump),
			"contract":     contractView,
			"yourSeat":     c.seat,
			"yourHand":     hand,
			"handCounts":   handCounts,
			"players":      playerNames,
			"currentTrick": trick,
			"tricksWon":    []int{countTricks(g, engine.TeamA), countTricks(g, engine.TeamB)},
			"scores":       []int{g.Teams[engine.TeamA].Score, g.Teams[engine.TeamB].Score},
			"legalMoves":   legalMoves,
			"biddingTurn":  g.Turn,
			"myAnnounces":  myAnnounces,
			"teamNames":    []string{r.TeamNames[0], r.TeamNames[1]},
		}
		b, _ := json.Marshal(msg)
		select {
		case c.send <- b:
		default:
		}
	}
}

func countTricks(g *engine.Game, team engine.TeamID) int {
	count := 0
	for _, trick := range g.TricksWon {
		// winner is whoever's card has max strength among trick, using
		// the same simple logic as engine (re-derive winner by seat).
		winnerSeat := trickWinnerSeat(g, trick)
		if engine.TeamOf(winnerSeat) == team {
			count++
		}
	}
	return count
}

func trickWinnerSeat(g *engine.Game, trick []engine.PlayedCard) int {
	winner := trick[0]
	for _, pc := range trick[1:] {
		if pc.Card.Suit == winner.Card.Suit {
			if g.CardStrength(pc.Card) > g.CardStrength(winner.Card) {
				winner = pc
			}
		} else if g.IsTrump(pc.Card.Suit) && !g.IsTrump(winner.Card.Suit) {
			winner = pc
		}
	}
	return winner.PlayerID
}

func (r *Room) broadcastTrickComplete(trick []engine.PlayedCard, winnerSeat int) {
	cards := make([]interface{}, len(trick))
	for i, pc := range trick {
		cards[i] = map[string]interface{}{"player": pc.PlayerID, "card": cardToStr(pc.Card)}
	}
	r.broadcastRaw(map[string]interface{}{
		"type":   "trick_complete",
		"winner": winnerSeat,
		"cards":  cards,
	})
}

func (r *Room) broadcastReadyStatus() {
	readyCount := 0
	total := 0
	for i := 0; i < 4; i++ {
		if r.Clients[i] != nil {
			total++
			if r.ReadyNext[i] {
				readyCount++
			}
		}
	}
	r.broadcastRaw(map[string]interface{}{
		"type":  "ready_status",
		"ready": readyCount,
		"total": total,
	})
}

func (r *Room) broadcastError(seat int, message string) {
	if seat < 0 || seat >= 4 || r.Clients[seat] == nil {
		return
	}
	b, _ := json.Marshal(map[string]interface{}{"type": "error", "message": message})
	select {
	case r.Clients[seat].send <- b:
	default:
	}
}

func (r *Room) broadcastHandResult(res engine.HandResult) {
	anns := make([]interface{}, 0, len(res.Announces.Announces))
	for _, a := range res.Announces.Announces {
		anns = append(anns, map[string]interface{}{
			"player": a.PlayerID,
			"value":  a.Value,
		})
	}
	r.broadcastRaw(map[string]interface{}{
		"type":           "hand_result",
		"trickPointsA":   res.TrickPoints[engine.TeamA],
		"trickPointsB":   res.TrickPoints[engine.TeamB],
		"contractMade":   res.ContractMade,
		"awardedTeam":    int(res.AwardedTeam),
		"awardedPoints":  res.AwardedPoints,
		"defenderPoints": res.DefenderPoints,
		"capo":           res.CapoFound,
		"belot":          res.BelotFound,
		"sequencePoints": res.Announces.SequencePoints,
		"carrePoints":    res.Announces.CarrePoints,
		"announces":      anns,
	})
}

func (r *Room) broadcastMatchOver(winner engine.TeamID) {
	r.broadcastRaw(map[string]interface{}{
		"type":   "match_over",
		"winner": int(winner),
	})
}
