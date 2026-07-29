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

// annKindToStr is a language-neutral wire token for the announce kind -
// the client looks up its own display text (in whichever language the
// player picked) from this instead of the server picking the wording.
func annKindToStr(k engine.AnnounceKind) string {
	switch k {
	case engine.AnnTierce:
		return "tierce"
	case engine.AnnFifty:
		return "fifty"
	case engine.AnnHundred:
		return "hundred"
	case engine.AnnCarreJ:
		return "carreJ"
	case engine.AnnCarreNine:
		return "carreNine"
	case engine.AnnCarreOther:
		return "carreOther"
	}
	return ""
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
