package engine

import "fmt"

// IsTrump reports whether the given suit counts as trump under the
// current contract (ContractAllTrump = every suit is trump,
// ContractNoTrump = no suit is trump, ContractSuit = only Game.Trump).
func (g *Game) IsTrump(s Suit) bool {
	if g.Contract == nil {
		return false
	}
	switch g.Contract.Type {
	case ContractAllTrump:
		return true
	case ContractNoTrump:
		return false
	default:
		return s == g.Trump
	}
}

// CardValue returns a card's point value under the current contract.
func (g *Game) CardValue(c Card) int {
	if g.IsTrump(c.Suit) {
		return c.Rank.TrumpValue()
	}
	return c.Rank.NonTrumpValue()
}

// CardStrength returns a card's relative strength under the current
// contract (only meaningful when comparing cards of the same suit).
func (g *Game) CardStrength(c Card) int {
	if g.IsTrump(c.Suit) {
		return trumpRankOrder[c.Rank]
	}
	return nonTrumpRankOrder[c.Rank]
}

// hasCard reports whether a hand contains the exact card.
func hasCard(hand []Card, c Card) bool {
	for _, hc := range hand {
		if hc == c {
			return true
		}
	}
	return false
}

// trickWinnerSoFar returns the seat currently winning the trick and the
// winning card, given the cards played so far (must be non-empty).
func (g *Game) trickWinnerSoFar(trick []PlayedCard) (int, Card) {
	ledSuit := trick[0].Card.Suit
	winner := trick[0]
	for _, pc := range trick[1:] {
		if pc.Card.Suit == winner.Card.Suit {
			if g.CardStrength(pc.Card) > g.CardStrength(winner.Card) {
				winner = pc
			}
		} else if g.IsTrump(pc.Card.Suit) && !g.IsTrump(winner.Card.Suit) {
			winner = pc
		}
		_ = ledSuit
	}
	return winner.PlayerID, winner.Card
}

// LegalMoves computes which cards a player may legally play right now.
//
// "Игра на козова боя" (plain suit contract): must follow the led suit
// if able. If you can't follow suit and the trick currently belongs to
// the OPPONENT, you must play trump (cut) if you hold any. If the
// opponent has already played trump in this trick, you must overtrump
// (play higher) if able - if you can't beat it, any trump you hold is
// fine. If your own partner is currently winning the trick, none of
// this cutting/overtrumping obligation applies - free discard. The
// same overtrump obligation also applies when the led suit IS trump
// itself (following suit into trump), again subject to the partner
// exception.
//
// "Без коз" (no trump): the only requirement is following suit - no
// trump exists, so no cut/overtrump obligation ever applies.
//
// "Всичко коз" (all-trump): must follow suit, and overtrumping is
// mandatory whenever you're able to, with NO partner exception -
// independent of who's currently winning. If you can't follow the led
// suit at all, a different suit can never win the trick under
// all-trump, so there's nothing to cut with - free discard.
func (g *Game) LegalMoves(playerID int) ([]Card, error) {
	if g.Phase != PhasePlaying {
		return nil, fmt.Errorf("not in playing phase")
	}
	if playerID != g.Turn {
		return nil, fmt.Errorf("not player %d's turn", playerID)
	}
	hand := g.Players[playerID].Hand

	if len(g.CurrentTrick) == 0 {
		// Leading the trick: any card is legal.
		return append([]Card{}, hand...), nil
	}

	contractType := ContractNone
	if g.Contract != nil {
		contractType = g.Contract.Type
	}

	ledSuit := g.CurrentTrick[0].Card.Suit
	partnerWinning := false
	if winnerID, _ := g.trickWinnerSoFar(g.CurrentTrick); TeamOf(winnerID) == TeamOf(playerID) {
		partnerWinning = true
	}
	// All-trump has no partner exception - overtrumping is always
	// mandatory there regardless of who's currently winning.
	partnerException := partnerWinning && contractType != ContractAllTrump

	following := filterBySuit(hand, ledSuit)
	if len(following) > 0 {
		if g.IsTrump(ledSuit) && !partnerException {
			_, winningCard := g.trickWinnerSoFar(g.CurrentTrick)
			higher := filterHigherTrump(following, winningCard, g)
			if len(higher) > 0 {
				return higher, nil
			}
		}
		return following, nil
	}

	// Cannot follow suit.
	if contractType == ContractAllTrump {
		// A different suit can never win the trick under all-trump, so
		// there's nothing to "cut" with - free discard.
		return append([]Card{}, hand...), nil
	}
	if contractType == ContractNoTrump {
		// No trump suit exists at all - free discard.
		return append([]Card{}, hand...), nil
	}
	if partnerException {
		return append([]Card{}, hand...), nil
	}

	trumpInHand := filterBySuit(hand, g.Trump)
	if len(trumpInHand) == 0 {
		return append([]Card{}, hand...), nil
	}

	trumpAlreadyPlayed := false
	for _, pc := range g.CurrentTrick {
		if g.IsTrump(pc.Card.Suit) {
			trumpAlreadyPlayed = true
			break
		}
	}
	if trumpAlreadyPlayed {
		_, winningCard := g.trickWinnerSoFar(g.CurrentTrick)
		higher := filterHigherTrump(trumpInHand, winningCard, g)
		if len(higher) > 0 {
			return higher, nil
		}
		// Have trump but none higher than the current winner - must
		// still play trump ("надцаква се, ако е възможно" - if not
		// possible, any trump is fine rather than a free discard).
		return trumpInHand, nil
	}

	// No trump played yet, partner not winning, and this player holds
	// trump: must cut.
	return trumpInHand, nil
}

func filterBySuit(hand []Card, s Suit) []Card {
	out := []Card{}
	for _, c := range hand {
		if c.Suit == s {
			out = append(out, c)
		}
	}
	return out
}

// filterHigherTrump returns cards from `cards` that are trump and beat
// `winningCard` in strength. If the contract is all-trump, any trump
// suit that's stronger counts (only same-suit strength is compared,
// consistent with normal trick rules - overtrumping always happens in
// the same suit as the winning trump card in standard belot).
func filterHigherTrump(cards []Card, winningCard Card, g *Game) []Card {
	out := []Card{}
	for _, c := range cards {
		if !g.IsTrump(c.Suit) {
			continue
		}
		if c.Suit == winningCard.Suit && g.CardStrength(c) > g.CardStrength(winningCard) {
			out = append(out, c)
		}
	}
	return out
}

// PlayCard validates and applies a card play. On completing the 8th
// card of a trick it resolves the winner, stores the trick, and sets up
// the next trick (or moves to scoring after 8 tricks).
func (g *Game) PlayCard(playerID int, c Card) error {
	legal, err := g.LegalMoves(playerID)
	if err != nil {
		return err
	}
	if !hasCard(g.Players[playerID].Hand, c) {
		return fmt.Errorf("player %d does not hold %v", playerID, c)
	}
	if !hasCard(legal, c) {
		return fmt.Errorf("illegal move: %v is not a legal play for player %d right now", c, playerID)
	}

	// Remove from hand.
	hand := g.Players[playerID].Hand
	for i, hc := range hand {
		if hc == c {
			g.Players[playerID].Hand = append(hand[:i], hand[i+1:]...)
			break
		}
	}

	g.CurrentTrick = append(g.CurrentTrick, PlayedCard{PlayerID: playerID, Card: c})

	if len(g.CurrentTrick) < 4 {
		g.Turn = (g.Turn + 1) % 4
		return nil
	}

	// Trick complete - resolve winner.
	winnerID, _ := g.trickWinnerSoFar(g.CurrentTrick)
	g.TricksWon = append(g.TricksWon, g.CurrentTrick)
	g.CurrentTrick = nil
	g.Turn = winnerID
	g.TrickLeader = winnerID

	if len(g.TricksWon) == 8 {
		g.Phase = PhaseScoring
	}
	return nil
}
