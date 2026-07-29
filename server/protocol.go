package server

import (
	"fmt"

	"belot/engine"
)

func suitToStr(s engine.Suit) string {
	switch s {
	case engine.Clubs:
		return "clubs"
	case engine.Diamonds:
		return "diamonds"
	case engine.Hearts:
		return "hearts"
	case engine.Spades:
		return "spades"
	}
	return ""
}

func strToSuit(s string) (engine.Suit, error) {
	switch s {
	case "clubs":
		return engine.Clubs, nil
	case "diamonds":
		return engine.Diamonds, nil
	case "hearts":
		return engine.Hearts, nil
	case "spades":
		return engine.Spades, nil
	}
	return 0, fmt.Errorf("unknown suit %q", s)
}

func rankToStr(r engine.Rank) string {
	return r.String()
}

// rankToBgDisplay renders a rank the way it should appear in a
// Bulgarian announce label - Асо/Поп/Дама/Вале for the face cards,
// but 7/8/9/10 stay as plain digits.
func rankToBgDisplay(r engine.Rank) string {
	switch r.String() {
	case "A":
		return "Асо"
	case "K":
		return "Поп"
	case "Q":
		return "Дама"
	case "J":
		return "Вале"
	default:
		return r.String()
	}
}

// suitToBgDisplay renders a suit the way it should appear in a
// Bulgarian announce label (capitalized, unlike suitToStr's lowercase
// wire-protocol token).
func suitToBgDisplay(s engine.Suit) string {
	switch s {
	case engine.Clubs:
		return "Спатия"
	case engine.Diamonds:
		return "Каро"
	case engine.Hearts:
		return "Купа"
	case engine.Spades:
		return "Пика"
	}
	return ""
}

func strToRank(s string) (engine.Rank, error) {
	for _, r := range engine.AllRanks {
		if r.String() == s {
			return r, nil
		}
	}
	return 0, fmt.Errorf("unknown rank %q", s)
}

func cardToStr(c engine.Card) string {
	return rankToStr(c.Rank) + "-" + suitToStr(c.Suit)
}

func strToCard(s string) (engine.Card, error) {
	// format "RANK-suit", e.g. "10-hearts", "J-clubs"
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '-' {
			r, err := strToRank(s[:i])
			if err != nil {
				return engine.Card{}, err
			}
			su, err := strToSuit(s[i+1:])
			if err != nil {
				return engine.Card{}, err
			}
			return engine.Card{Suit: su, Rank: r}, nil
		}
	}
	return engine.Card{}, fmt.Errorf("bad card format %q", s)
}

func contractTypeToStr(ct engine.ContractType) string {
	switch ct {
	case engine.ContractSuit:
		return "suit"
	case engine.ContractNoTrump:
		return "notrump"
	case engine.ContractAllTrump:
		return "alltrump"
	}
	return ""
}

func describeAnnounce(a engine.Announce) string {
	switch a.Kind {
	case engine.AnnTierce:
		return fmt.Sprintf("Терца до %s %s (20)", rankToBgDisplay(a.HighRank), suitToBgDisplay(a.Suit))
	case engine.AnnFifty:
		return fmt.Sprintf("Петдесет до %s %s (50)", rankToBgDisplay(a.HighRank), suitToBgDisplay(a.Suit))
	case engine.AnnHundred:
		return fmt.Sprintf("Сто до %s %s (100)", rankToBgDisplay(a.HighRank), suitToBgDisplay(a.Suit))
	case engine.AnnCarreJ:
		return "Каре валета (200)"
	case engine.AnnCarreNine:
		return "Каре девятки (150)"
	case engine.AnnCarreOther:
		return fmt.Sprintf("Каре %s (100)", rankToBgDisplay(a.HighRank))
	}
	return fmt.Sprintf("Анонс (%d)", a.Value)
}

func phaseToStr(p engine.GamePhase) string {
	switch p {
	case engine.PhaseDealing:
		return "dealing"
	case engine.PhaseBidding:
		return "bidding"
	case engine.PhaseAnnouncing:
		return "announcing"
	case engine.PhasePlaying:
		return "playing"
	case engine.PhaseScoring:
		return "scoring"
	case engine.PhaseFinished:
		return "finished"
	}
	return ""
}
