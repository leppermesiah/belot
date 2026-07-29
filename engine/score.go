package engine

// CapoBonus (валат) is the extra flat premium awarded when one team
// takes all 8 tricks. Common table value - adjust here if your house
// rules use a different number.
const CapoBonus = 90

// HandResult is the full breakdown of one completed hand's scoring.
type HandResult struct {
	TrickPoints    map[TeamID]int // after any "без коз" doubling and capo bonus
	LastTrickTeam  TeamID
	Announces      AnnounceResult
	BelotTeam      TeamID
	BelotFound     bool
	CapoTeam       TeamID
	CapoFound      bool // "валат" - one team took all 8 tricks
	ContractMade   bool
	AwardedTeam    TeamID // team that ends up scoring this hand's pot (contract success) or the whole pot (вътре)
	AwardedPoints  int    // final points added to AwardedTeam's match score (post rounding + contra/reconto)
	DefenderPoints int    // points the OTHER team records this hand (0 on вътре)
}

// DeclaredAnnounces tracks which teams actually clicked "declare" for
// a sequence, a carre, and/or belot during this hand - per the real
// table rule, an undeclared combo (even if it's genuinely the best
// one in someone's hand) does NOT count toward scoring. Compare
// against a nil pointer in ScoreHand, which scores everything
// automatically regardless of declaration (kept for existing tests
// and any caller that doesn't care about the declare-to-count rule).
type DeclaredAnnounces struct {
	Sequence map[TeamID]bool
	Carre    map[TeamID]bool
	Belot    map[TeamID]bool
}

func (d *DeclaredAnnounces) sequenceDeclared(team TeamID) bool {
	if d == nil {
		return true
	}
	return d.Sequence[team]
}
func (d *DeclaredAnnounces) carreDeclared(team TeamID) bool {
	if d == nil {
		return true
	}
	return d.Carre[team]
}
func (d *DeclaredAnnounces) belotDeclared(team TeamID) bool {
	if d == nil {
		return true
	}
	return d.Belot[team]
}

// ScoreHand computes the full result of a completed hand exactly as
// scoreHand does, but treats every detected sequence/carre/belot as
// already declared (auto-scores everything). This is the pre-existing
// behavior kept for callers (mostly tests) that don't model the
// declare-to-count rule explicitly.
func ScoreHand(g *Game) HandResult {
	return scoreHand(g, nil)
}

// ScoreHandWithDeclared is the real production path: only a
// sequence/carre/belot the relevant team actually clicked "declare"
// for during play counts toward scoring - matching the real-table
// rule that an unannounced combo is forfeited even if it was
// genuinely the better hand.
func ScoreHandWithDeclared(g *Game, declared DeclaredAnnounces) HandResult {
	return scoreHand(g, &declared)
}

// ScoreHand computes the full result of a completed hand (Phase must be
// PhaseScoring, i.e. all 8 tricks played) and adds the outcome to each
// team's running match score. Implements:
//   - card point totals (doubled under "без коз")
//   - last-trick +10, валат/капо premium
//   - announces compared separately by category (sequences vs carres),
//     only credited to a team that actually declared that category
//   - белот, only credited if actually declared
//   - "вътре": if the calling team doesn't outscore the defenders, the
//     calling team gets nothing and the whole pot goes to the defenders
//   - contra/reconto multiplier (also applied to all premiums)
//   - closest-10 rounding, with the official tie-break for the
//     "self-paired last digit" edge case (smaller raw total rounds up,
//     larger rounds down)
//
// NOT implemented: rule 4 "висяща игра" (the carry-over/hanging-points
// mechanic for an exact raw point tie between the two teams). That's a
// genuinely rare edge case (a perfect 50/50 raw split) and the source
// text was ambiguous enough about the exact carry-over bookkeeping that
// guessing wrong seemed worse than flagging it. On an exact raw tie,
// this implementation currently just falls back to normal "вътре"
// (whole pot to the defending team) - deliberately, so it's confirmed-
// safe rather than confidently wrong. If you hit this in a real game,
// let me know the outcome you'd expect and I'll wire it up properly.
func scoreHand(g *Game, declared *DeclaredAnnounces) HandResult {
	res := HandResult{TrickPoints: map[TeamID]int{TeamA: 0, TeamB: 0}}

	// 1. Raw trick points.
	tricksWonBy := map[TeamID]int{}
	for i, trick := range g.TricksWon {
		winnerID, _ := g.trickWinnerSoFar(trick)
		team := TeamOf(winnerID)
		tricksWonBy[team]++
		for _, pc := range trick {
			res.TrickPoints[team] += g.CardValue(pc.Card)
		}
		if i == len(g.TricksWon)-1 {
			res.LastTrickTeam = team
			res.TrickPoints[team] += 10
		}
	}

	// 2. "Без коз" scores double on the raw card+last-trick points
	// (общо 130 -> 260, per the official rounding table). This does
	// NOT double capo/announces/belot - those are flat premiums added
	// after this step.
	if g.Contract != nil && g.Contract.Type == ContractNoTrump {
		res.TrickPoints[TeamA] *= 2
		res.TrickPoints[TeamB] *= 2
	}

	// 3. Capo/валат - one team took all 8 tricks.
	if tricksWonBy[TeamA] == 8 {
		res.CapoFound, res.CapoTeam = true, TeamA
		res.TrickPoints[TeamA] += CapoBonus
	} else if tricksWonBy[TeamB] == 8 {
		res.CapoFound, res.CapoTeam = true, TeamB
		res.TrickPoints[TeamB] += CapoBonus
	}

	// 4. Announces - sequences and carres compared separately; "без
	// коз" allows none at all (handled inside CompareAnnounces). Only
	// credited to whichever team actually declared that category.
	res.Announces = CompareAnnounces(g)
	sequenceCounts := res.Announces.SequenceFound && declared.sequenceDeclared(res.Announces.SequenceTeam)
	carreCounts := res.Announces.CarreFound && declared.carreDeclared(res.Announces.CarreTeam)

	// 5. Belot - always scores for its holder regardless of the above,
	// never under "без коз" - but only if actually declared.
	belotFound, belotTeam, _, _ := DetectBelot(g, g.OriginalHands)
	res.BelotFound = belotFound && declared.belotDeclared(belotTeam)
	if res.BelotFound {
		res.BelotTeam = belotTeam
	}

	callingTeam := g.Contract.CallerTeam
	defendingTeam := TeamA
	if callingTeam == TeamA {
		defendingTeam = TeamB
	}

	callingTotal := res.TrickPoints[callingTeam]
	defendingTotal := res.TrickPoints[defendingTeam]
	if sequenceCounts {
		if res.Announces.SequenceTeam == callingTeam {
			callingTotal += res.Announces.SequencePoints
		} else {
			defendingTotal += res.Announces.SequencePoints
		}
	}
	if carreCounts {
		if res.Announces.CarreTeam == callingTeam {
			callingTotal += res.Announces.CarrePoints
		} else {
			defendingTotal += res.Announces.CarrePoints
		}
	}
	if res.BelotFound {
		if res.BelotTeam == callingTeam {
			callingTotal += 20
		} else {
			defendingTotal += 20
		}
	}

	mult := g.Contract.Multiplier()

	// 6. Contract success check: calling team must score STRICTLY more
	// than the defending team. If not, it's "вътре" - the calling team
	// gets nothing and the ENTIRE pot goes to the defenders.
	if callingTotal > defendingTotal {
		res.ContractMade = true
		roundedCalling, roundedDefending := roundHandScore(callingTotal, defendingTotal, g.Contract.Type)
		res.AwardedTeam = callingTeam
		res.AwardedPoints = roundedCalling * mult
		res.DefenderPoints = roundedDefending
		g.Teams[defendingTeam].Score += res.DefenderPoints / 10
	} else {
		res.ContractMade = false
		res.AwardedTeam = defendingTeam
		pot := callingTotal + defendingTotal
		roundedPot := roundToNearest10(pot)
		res.AwardedPoints = roundedPot * mult
		res.DefenderPoints = 0
	}

	// The match's running score (compared against TargetScore, 151) is
	// tracked in the "divided by 10" units belot.bg's own scoresheet
	// convention uses - AwardedPoints/DefenderPoints stay in full raw
	// units on the HandResult itself (useful for display/logging), but
	// only the /10 value actually accumulates toward winning the match.
	g.Teams[res.AwardedTeam].Score += res.AwardedPoints / 10
	g.Phase = PhaseFinished
	return res
}

// roundToNearest10 uses standard round-half-up.
func roundToNearest10(x int) int {
	rem := x % 10
	if rem >= 5 {
		return x - rem + 10
	}
	return x - rem
}

// problemDigit is the official "self-paired last digit" tie-break
// trigger, which depends on the contract type:
//   - всичко коз (base pool 258): digit 4
//   - без коз (base pool 260, already doubled): digit 5
//   - suit contract / игра на коз (base pool 162): digit 6
//
// This is NOT "any matching last digit" - it's specifically THIS one
// digit per contract type, since that's the only digit where the pool
// total's own last digit forces both teams' totals to a genuine
// halfway-rounding ambiguity. A coincidental match at some OTHER
// digit (e.g. both totals happening to end in 7) is not a real tie
// and each side just rounds independently as normal.
func problemDigit(contractType ContractType) int {
	switch contractType {
	case ContractAllTrump:
		return 4
	case ContractNoTrump:
		return 5
	default: // suit contract
		return 6
	}
}

// roundHandScore rounds both teams' point totals to the nearest 10 for
// recording, honoring the official tie-break rule: when both totals
// happen to end in that contract type's specific "problem digit" (the
// self-paired edge case where standard rounding would send both the
// same direction and break the total), the team with FEWER raw points
// rounds UP to the next 10, and the team with MORE raw points rounds
// DOWN (floor) to the same decade.
func roundHandScore(callingTotal, defendingTotal int, contractType ContractType) (int, int) {
	digit := problemDigit(contractType)
	dCalling, dDefending := callingTotal%10, defendingTotal%10
	if dCalling == digit && dDefending == digit {
		if callingTotal < defendingTotal {
			return callingTotal - dCalling + 10, defendingTotal - dDefending
		}
		return callingTotal - dCalling, defendingTotal - dDefending + 10
	}
	return roundToNearest10(callingTotal), roundToNearest10(defendingTotal)
}
