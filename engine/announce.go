package engine

import "sort"

type AnnounceCategory int

const (
	CategorySequence AnnounceCategory = iota // терца/петдесет/сто
	CategoryCarre
)

type AnnounceKind int

const (
	AnnTierce     AnnounceKind = iota // 3 in sequence
	AnnFifty                          // 4 in sequence
	AnnHundred                        // 5+ in sequence
	AnnCarreJ                         // four jacks
	AnnCarreNine                      // four nines
	AnnCarreOther                     // four aces/tens/kings/queens
)

type Announce struct {
	PlayerID int
	Category AnnounceCategory
	Kind     AnnounceKind
	Suit     Suit // sequence suit, or arbitrary for carre
	HighRank Rank // top card of the sequence, or the carre's rank
	Value    int
	Cards    []Card // the actual cards this announce is built from
}

// sequenceRankOrder is the fixed natural order used for straights,
// independent of the trump contract.
var sequenceRankOrder = map[Rank]int{
	Seven: 0, Eight: 1, Nine: 2, Ten: 3, Jack: 4, Queen: 5, King: 6, Ace: 7,
}

// detectSequencesInSuit finds all non-overlapping runs of length >= 3
// within a single suit for one player's hand.
func detectSequencesInSuit(suit Suit, cards []Card) []Announce {
	ranks := make([]Rank, len(cards))
	for i, c := range cards {
		ranks[i] = c.Rank
	}
	if len(ranks) < 3 {
		return nil
	}
	sort.Slice(ranks, func(i, j int) bool {
		return sequenceRankOrder[ranks[i]] < sequenceRankOrder[ranks[j]]
	})
	var out []Announce
	runStart := 0
	for i := 1; i <= len(ranks); i++ {
		broken := i == len(ranks) || sequenceRankOrder[ranks[i]] != sequenceRankOrder[ranks[i-1]]+1
		if broken {
			runLen := i - runStart
			if runLen >= 3 {
				kind := AnnHundred
				val := 100
				switch runLen {
				case 3:
					kind, val = AnnTierce, 20
				case 4:
					kind, val = AnnFifty, 50
				}
				runCards := make([]Card, runLen)
				for j := 0; j < runLen; j++ {
					runCards[j] = Card{Suit: suit, Rank: ranks[runStart+j]}
				}
				out = append(out, Announce{
					Category: CategorySequence,
					Kind:     kind,
					Suit:     suit,
					HighRank: ranks[i-1],
					Value:    val,
					Cards:    runCards,
				})
			}
			runStart = i
		}
	}
	return out
}

// DetectAnnounces returns every sequence/carre announce a player's hand
// supports, with rule 12 overlap resolution already applied (if a card
// is shared between a carre and a sequence, only the higher-value one
// is kept). Does not include belot (see DetectBelot) and returns
// nothing under a "без коз" contract, where no combinations may be
// declared at all.
func DetectAnnounces(g *Game, playerID int, hand []Card) []Announce {
	if g.Contract != nil && g.Contract.Type == ContractNoTrump {
		return nil
	}

	var sequences []Announce
	var carres []Announce

	bySuit := map[Suit][]Card{}
	byRank := map[Rank][]Card{}
	for _, c := range hand {
		bySuit[c.Suit] = append(bySuit[c.Suit], c)
		byRank[c.Rank] = append(byRank[c.Rank], c)
	}

	for suit, cards := range bySuit {
		for _, ann := range detectSequencesInSuit(suit, cards) {
			ann.PlayerID = playerID
			sequences = append(sequences, ann)
		}
	}

	// Carres: four of a kind. Only J, 9, A, 10, K, Q count - not 7/8.
	for rank, cards := range byRank {
		if len(cards) != 4 {
			continue
		}
		var kind AnnounceKind
		var val int
		switch rank {
		case Jack:
			kind, val = AnnCarreJ, 200
		case Nine:
			kind, val = AnnCarreNine, 150
		case Ace, Ten, King, Queen:
			kind, val = AnnCarreOther, 100
		default:
			continue // 7s and 8s don't count as a carre
		}
		carres = append(carres, Announce{
			PlayerID: playerID,
			Category: CategoryCarre,
			Kind:     kind,
			HighRank: rank,
			Value:    val,
			Cards:    append([]Card{}, cards...),
		})
	}

	resolveCarreSequenceOverlap(&sequences, &carres)

	out := make([]Announce, 0, len(sequences)+len(carres))
	out = append(out, sequences...)
	out = append(out, carres...)
	return out
}

// resolveCarreSequenceOverlap implements rule 12: if a single physical
// card is part of both a carre and a sequence, the player can only
// declare one of the two - we resolve this automatically in favor of
// whichever is worth more, dropping the lesser one entirely.
func resolveCarreSequenceOverlap(sequences, carres *[]Announce) {
	shareCard := func(a, b Announce) bool {
		for _, ca := range a.Cards {
			for _, cb := range b.Cards {
				if ca == cb {
					return true
				}
			}
		}
		return false
	}

	var keptSeq []Announce
	droppedCarre := map[int]bool{} // index into *carres
	for _, s := range *sequences {
		conflict := false
		for ci, c := range *carres {
			if droppedCarre[ci] {
				continue
			}
			if shareCard(s, c) {
				conflict = true
				if s.Value <= c.Value {
					// sequence loses, drop it, keep the carre
				} else {
					droppedCarre[ci] = true
				}
			}
		}
		if !conflict {
			keptSeq = append(keptSeq, s)
			continue
		}
		// Re-check: keep the sequence only if it beat every carre it
		// conflicted with (i.e. none of those carres survived).
		stillConflicting := false
		for ci, c := range *carres {
			if !droppedCarre[ci] && shareCard(s, c) {
				stillConflicting = true
			}
		}
		if !stillConflicting {
			keptSeq = append(keptSeq, s)
		}
	}

	var keptCarre []Announce
	for ci, c := range *carres {
		if !droppedCarre[ci] {
			keptCarre = append(keptCarre, c)
		}
	}

	*sequences = keptSeq
	*carres = keptCarre
}

// bestInCategory returns the strongest single announce from a list
// (assumed to be pre-filtered to one category), using value first, then
// highest card rank as tie-break.
func bestInCategory(anns []Announce) (Announce, bool) {
	if len(anns) == 0 {
		return Announce{}, false
	}
	best := anns[0]
	for _, a := range anns[1:] {
		if a.Value > best.Value {
			best = a
			continue
		}
		if a.Value == best.Value && sequenceRankOrder[a.HighRank] > sequenceRankOrder[best.HighRank] {
			best = a
		}
	}
	return best, true
}

// AnnounceResult is the outcome of comparing both announce categories
// across all 4 players for one hand.
type AnnounceResult struct {
	SequenceFound  bool
	SequenceTeam   TeamID
	SequencePoints int
	CarreFound     bool
	CarreTeam      TeamID
	CarrePoints    int
	Announces      []Announce // every announce that actually scored
}

// CompareAnnounces implements rule 13: sequences (терца/петдесет/сто)
// and carres are compared SEPARATELY. In each category, only the team
// holding the single best announce scores ANY points in that category -
// the opposing team's announces in that category are wiped out, even if
// they had some. If the two best announces are EXACTLY tied (same value
// and same high card), every announce in that category is voided for
// both teams. A hand can have different winners for each category.
func CompareAnnounces(g *Game) AnnounceResult {
	var res AnnounceResult

	perPlayer := make([][]Announce, 4)
	for i := 0; i < 4; i++ {
		perPlayer[i] = DetectAnnounces(g, i, g.OriginalHands[i])
	}

	byCategory := func(cat AnnounceCategory) [4][]Announce {
		var out [4][]Announce
		for i := 0; i < 4; i++ {
			for _, a := range perPlayer[i] {
				if a.Category == cat {
					out[i] = append(out[i], a)
				}
			}
		}
		return out
	}

	compare := func(perPlayerCat [4][]Announce) (team TeamID, points int, found bool, all []Announce) {
		var teamBest [2]Announce
		var teamHasBest [2]bool
		for i := 0; i < 4; i++ {
			b, ok := bestInCategory(perPlayerCat[i])
			if !ok {
				continue
			}
			team := TeamOf(i)
			if !teamHasBest[team] || b.Value > teamBest[team].Value ||
				(b.Value == teamBest[team].Value && sequenceRankOrder[b.HighRank] > sequenceRankOrder[teamBest[team].HighRank]) {
				teamBest[team] = b
				teamHasBest[team] = true
			}
		}

		if !teamHasBest[TeamA] && !teamHasBest[TeamB] {
			return TeamA, 0, false, nil
		}
		var winner TeamID
		switch {
		case teamHasBest[TeamA] && !teamHasBest[TeamB]:
			winner = TeamA
		case teamHasBest[TeamB] && !teamHasBest[TeamA]:
			winner = TeamB
		default:
			a, b := teamBest[TeamA], teamBest[TeamB]
			if a.Value == b.Value && a.HighRank == b.HighRank {
				return TeamA, 0, false, nil // exact tie - voided for both
			}
			if a.Value > b.Value || (a.Value == b.Value && sequenceRankOrder[a.HighRank] > sequenceRankOrder[b.HighRank]) {
				winner = TeamA
			} else {
				winner = TeamB
			}
		}

		total := 0
		var scored []Announce
		for i := 0; i < 4; i++ {
			if TeamOf(i) != winner {
				continue
			}
			for _, a := range perPlayerCat[i] {
				total += a.Value
				scored = append(scored, a)
			}
		}
		return winner, total, true, scored
	}

	seqTeam, seqPoints, seqFound, seqAnns := compare(byCategory(CategorySequence))
	res.SequenceFound, res.SequenceTeam, res.SequencePoints = seqFound, seqTeam, seqPoints
	res.Announces = append(res.Announces, seqAnns...)

	carreTeam, carrePoints, carreFound, carreAnns := compare(byCategory(CategoryCarre))
	res.CarreFound, res.CarreTeam, res.CarrePoints = carreFound, carreTeam, carrePoints
	res.Announces = append(res.Announces, carreAnns...)

	return res
}

// DetectBelot checks each player's ORIGINAL 8-card hand for K+Q of the
// trump suit. Belot always scores for its holder's team regardless of
// the announce comparison above. Only applies when a trump suit exists
// (ContractSuit or ContractAllTrump - "всичко коз" has trump on every
// suit, so belot can technically apply in whichever suit the pair sits
// in, per common house rules). Never applies under "без коз" (no trump
// suit exists at all).
// DetectBelot finds a K+Q-of-trump pair in someone's original hand.
// Under a suit contract or без коз... wait, без коз never gets here
// (guarded above) - under a plain suit contract, holding K+Q of the
// declared trump suit always counts, played whenever. Under всичко
// коз specifically, every suit is technically trump, but belot only
// actually counts if the King or Queen was played either as the
// trick's leading card, or while that suit happened to be the one
// currently led - an off-suit forced discard of K/Q (because the
// player was out of the led suit) does NOT count, even though that
// suit is nominally trump too.
func DetectBelot(g *Game, originalHands [4][]Card) (found bool, team TeamID, playerID int, suit Suit) {
	if g.Contract == nil || g.Contract.Type == ContractNoTrump {
		return false, 0, 0, 0
	}
	allTrump := g.Contract.Type == ContractAllTrump
	for i := 0; i < 4; i++ {
		bySuit := map[Suit]map[Rank]bool{}
		for _, c := range originalHands[i] {
			if bySuit[c.Suit] == nil {
				bySuit[c.Suit] = map[Rank]bool{}
			}
			bySuit[c.Suit][c.Rank] = true
		}
		for s, ranks := range bySuit {
			if !ranks[King] || !ranks[Queen] || !g.IsTrump(s) {
				continue
			}
			if allTrump && !belotPlayedValidlyUnderAllTrump(g, i, s) {
				continue
			}
			return true, TeamOf(i), i, s
		}
	}
	return false, 0, 0, 0
}

// belotPlayedValidlyUnderAllTrump checks whether player i actually
// played their King or Queen of suit s as the leading card of some
// trick, or while suit s was the one currently led.
func belotPlayedValidlyUnderAllTrump(g *Game, playerID int, s Suit) bool {
	for _, trick := range g.TricksWon {
		if len(trick) == 0 {
			continue
		}
		ledSuit := trick[0].Card.Suit
		for pos, pc := range trick {
			if pc.PlayerID != playerID || pc.Card.Suit != s {
				continue
			}
			if pc.Card.Rank != King && pc.Card.Rank != Queen {
				continue
			}
			if pos == 0 || ledSuit == s {
				return true
			}
		}
	}
	return false
}
