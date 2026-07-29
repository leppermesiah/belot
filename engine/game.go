package engine

import (
	"math/rand"
)

// Player is a seat at the table. 4 players, seats 0-3, alternating teams
// (0&2 vs 1&3), turn order goes clockwise (0->1->2->3->0).
type Player struct {
	ID   int
	Name string
	Hand []Card
}

// TeamID identifies one of the two teams.
type TeamID int

const (
	TeamA TeamID = iota // players 0, 2
	TeamB               // players 1, 3
)

// Team tracks the running score for one of the two teams.
type Team struct {
	ID    TeamID
	Score int
}

// TeamOf returns which team a given seat/player index belongs to.
func TeamOf(playerID int) TeamID {
	if playerID%2 == 0 {
		return TeamA
	}
	return TeamB
}

// GamePhase tracks which stage of a single hand we're in.
type GamePhase int

const (
	PhaseDealing GamePhase = iota
	PhaseBidding
	PhaseAnnouncing // белот / anonси declared on first card of the trick they apply to
	PhasePlaying
	PhaseScoring
	PhaseFinished
)

// Game holds the full state for one hand (deal) of Belot.
type Game struct {
	Players       [4]*Player
	Teams         [2]*Team
	Dealer        int // seat index of the current dealer
	Trump         Suit
	TrumpChosen   bool
	Contract      *Contract
	Phase         GamePhase
	CurrentTrick  []PlayedCard
	TrickLeader   int            // seat that leads the current trick
	Turn          int            // seat whose turn it is to act
	TricksWon     [][]PlayedCard // history of completed tricks this hand
	OriginalHands [4][]Card      // snapshot of each FULL hand once all 8 cards are dealt (for announces/belot)
	pendingKitty  [4][]Card      // each player's remaining 3 cards, held back until bidding resolves
	rng           *rand.Rand
}

// PlayedCard records a card played to a trick along with who played it.
type PlayedCard struct {
	PlayerID int
	Card     Card
}

// NewGame creates a fresh game with 4 named players. dealer is the seat
// index (0-3) that deals first.
func NewGame(names [4]string, dealer int, seed int64) *Game {
	g := &Game{
		Dealer: dealer,
		Phase:  PhaseDealing,
		rng:    rand.New(rand.NewSource(seed)),
	}
	for i := 0; i < 4; i++ {
		g.Players[i] = &Player{ID: i, Name: names[i]}
	}
	g.Teams[TeamA] = &Team{ID: TeamA}
	g.Teams[TeamB] = &Team{ID: TeamB}
	return g
}

// Deal shuffles a fresh 32-card deck and deals the first 5 cards to
// each player - bidding happens on this partial hand, per the table's
// house rule. The remaining 3 cards per player are held back and only
// dealt once bidding resolves (see DealRemaining).
func (g *Game) Deal() {
	deck := NewDeck()
	g.rng.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	for i := 0; i < 4; i++ {
		g.Players[i].Hand = append([]Card{}, deck[i*5:(i+1)*5]...)
		g.pendingKitty[i] = append([]Card{}, deck[20+i*3:20+(i+1)*3]...)
	}
	g.Phase = PhaseBidding
	// Bidding starts with the player after the dealer.
	g.Turn = (g.Dealer + 1) % 4
}

// DealRemaining hands out the held-back 3 cards per player once bidding
// has resolved into a contract, completing each hand to 8 cards, and
// snapshots the now-full hands for later announce/belot detection.
func (g *Game) DealRemaining() {
	for i := 0; i < 4; i++ {
		g.Players[i].Hand = append(g.Players[i].Hand, g.pendingKitty[i]...)
		g.pendingKitty[i] = nil
		g.OriginalHands[i] = append([]Card{}, g.Players[i].Hand...)
	}
}
