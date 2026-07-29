package engine

import "fmt"

// ContractType is what a bid names: a suit trump, no-trump, or all-trump.
type ContractType int

const (
	ContractNone ContractType = iota // no contract yet / passed
	ContractSuit
	ContractNoTrump
	ContractAllTrump
)

// Contract is the current (or final) bid of a hand.
type Contract struct {
	Type       ContractType
	Suit       Suit // meaningful only if Type == ContractSuit
	CallerID   int  // seat that made this call
	CallerTeam TeamID
	Contra     bool // opponents doubled this call
	Reconto    bool // calling team redoubled after contra
}

// Multiplier returns the score multiplier from contra/reconto.
func (c Contract) Multiplier() int {
	if c.Reconto {
		return 4
	}
	if c.Contra {
		return 2
	}
	return 1
}

// contractRank gives every possible call a position in the standard
// Bulgarian belot auction order, weakest to strongest:
//
//	♣ спатия < ♦ каро < ♥ купа < ♠ пика < без коз < всичко коз
//
// Any new call must outrank the current best call - this is what lets a
// contra'd team "escape" by bidding higher instead of accepting the
// double, and it's why "всичко коз" has nowhere higher left to escape
// to (it's the top of the order).
func contractRank(ct ContractType, suit Suit) int {
	switch ct {
	case ContractSuit:
		switch suit {
		case Clubs:
			return 0
		case Diamonds:
			return 1
		case Hearts:
			return 2
		case Spades:
			return 3
		}
	case ContractNoTrump:
		return 4
	case ContractAllTrump:
		return 5
	}
	return -1
}

// BiddingState runs a real auction: any player, on their turn, may pass,
// outbid the current best call with something strictly higher in
// contractRank, contra an opponent's live call, or (if contra'd)
// reconto their own team's call. Bidding closes after 3 consecutive
// passes following the last call/contra/reconto action. If all 4
// players pass before anyone ever calls, the hand is redealt.
type BiddingState struct {
	Game              *Game
	BestCall          *Contract // nil until someone calls
	ConsecutivePasses int       // passes in a row since the last action
	PassesWithNoCall  int       // for the "everyone passed, redeal" case
	Done              bool
}

func NewBidding(g *Game) *BiddingState {
	return &BiddingState{Game: g}
}

// CallSuit names a suit as trump.
func (b *BiddingState) CallSuit(playerID int, suit Suit) error {
	return b.call(playerID, ContractSuit, suit)
}

// CallNoTrump names a "no trump" contract (без коз).
func (b *BiddingState) CallNoTrump(playerID int) error {
	return b.call(playerID, ContractNoTrump, 0)
}

// CallAllTrump names an "all trump" contract (всичко коз).
func (b *BiddingState) CallAllTrump(playerID int) error {
	return b.call(playerID, ContractAllTrump, 0)
}

func (b *BiddingState) call(playerID int, ct ContractType, suit Suit) error {
	if err := b.checkTurn(playerID); err != nil {
		return err
	}
	newRank := contractRank(ct, suit)
	if b.BestCall != nil && newRank <= contractRank(b.BestCall.Type, b.BestCall.Suit) {
		return fmt.Errorf("must call something higher than the current bid")
	}
	b.BestCall = &Contract{
		Type:       ct,
		Suit:       suit,
		CallerID:   playerID,
		CallerTeam: TeamOf(playerID),
	}
	// Publish onto the Game immediately (same pointer, so later
	// Contra/Reconto mutations are visible automatically) so every
	// player sees the live contract right away, not only once the
	// whole auction finishes.
	b.Game.Contract = b.BestCall
	if ct == ContractSuit {
		b.Game.Trump = suit
	}
	b.Game.TrumpChosen = true
	b.ConsecutivePasses = 0
	b.Game.Turn = (b.Game.Turn + 1) % 4
	return nil
}

// Pass records a pass from the current turn player.
func (b *BiddingState) Pass(playerID int) error {
	if err := b.checkTurn(playerID); err != nil {
		return err
	}
	if b.BestCall == nil {
		b.PassesWithNoCall++
		if b.PassesWithNoCall >= 4 {
			b.Done = true // nobody ever called - redeal
			return nil
		}
		b.Game.Turn = (b.Game.Turn + 1) % 4
		return nil
	}
	b.ConsecutivePasses++
	b.Game.Turn = (b.Game.Turn + 1) % 4
	if b.ConsecutivePasses >= 3 {
		b.finish()
	}
	return nil
}

// Contra lets an opponent of the current caller double the live call.
// Doubling can still be "escaped" afterwards if the calling side bids
// something higher instead of passing/reconto-ing.
func (b *BiddingState) Contra(playerID int) error {
	if err := b.checkTurn(playerID); err != nil {
		return err
	}
	if b.BestCall == nil {
		return fmt.Errorf("no contract to contra")
	}
	if TeamOf(playerID) == b.BestCall.CallerTeam {
		return fmt.Errorf("cannot contra your own team's contract")
	}
	if b.BestCall.Contra {
		return fmt.Errorf("already contra'd")
	}
	b.BestCall.Contra = true
	b.ConsecutivePasses = 0
	b.Game.Turn = (b.Game.Turn + 1) % 4
	return nil
}

// Reconto lets the calling team redouble after being contra'd. On
// "всичко коз" this is the last possible escalation - there's no
// higher call left to escape to.
func (b *BiddingState) Reconto(playerID int) error {
	if err := b.checkTurn(playerID); err != nil {
		return err
	}
	if b.BestCall == nil || !b.BestCall.Contra {
		return fmt.Errorf("no contra to reconto")
	}
	if TeamOf(playerID) != b.BestCall.CallerTeam {
		return fmt.Errorf("only the calling team may reconto")
	}
	if b.BestCall.Reconto {
		return fmt.Errorf("already reconto'd")
	}
	b.BestCall.Reconto = true
	b.ConsecutivePasses = 0
	b.Game.Turn = (b.Game.Turn + 1) % 4
	return nil
}

func (b *BiddingState) finish() {
	b.Done = true
	b.Game.DealRemaining()
	b.Game.Phase = PhasePlaying
	// Whoever is left of the dealer leads the first trick.
	b.Game.Turn = (b.Game.Dealer + 1) % 4
	b.Game.TrickLeader = b.Game.Turn
}

func (b *BiddingState) checkTurn(playerID int) error {
	if b.Done {
		return fmt.Errorf("bidding already finished")
	}
	if playerID != b.Game.Turn {
		return fmt.Errorf("not player %d's turn (turn is seat %d)", playerID, b.Game.Turn)
	}
	return nil
}
