package engine

import (
	"crypto/rand"
	"encoding/binary"
	"time"
)

// Match plays repeated hands with dealer rotation until one team reaches
// TargetScore. This is the top-level object the server drives.
type Match struct {
	Names       [4]string
	TargetScore int
	Dealer      int
	TeamScores  [2]int
	HandNumber  int
	History     []HandResult
	Current     *Game
	Bidding     *BiddingState
}

func NewMatch(names [4]string, targetScore int) *Match {
	if targetScore <= 0 {
		targetScore = 151 // common Bulgarian belot target
	}
	return &Match{Names: names, TargetScore: targetScore, Dealer: 0}
}

// freshShuffleSeed mixes wall-clock time with crypto/rand entropy, so
// the deal is unpredictable regardless of how the process was started
// (this belt-and-suspenders approach removes any doubt, on top of
// simply not reusing a fixed per-match counter).
func freshShuffleSeed() int64 {
	var buf [8]byte
	seed := time.Now().UnixNano()
	if _, err := rand.Read(buf[:]); err == nil {
		seed ^= int64(binary.LittleEndian.Uint64(buf[:]))
	}
	return seed
}

// StartHand deals a new hand and opens bidding. If all 4 players pass
// with no contract, call StartHand again (it advances the dealer for a
// clean redeal automatically when Current is nil'd out by RedealIfNoBid).
func (m *Match) StartHand() {
	m.HandNumber++
	seed := freshShuffleSeed()
	g := NewGame(m.Names, m.Dealer, seed)
	g.Teams[TeamA].Score = m.TeamScores[TeamA]
	g.Teams[TeamB].Score = m.TeamScores[TeamB]
	g.Deal()
	m.Current = g
	m.Bidding = NewBidding(g)
}

// AfterBiddingResolved should be called once m.Bidding.Done is true. If
// nobody ever called (BestCall is nil), it redeals with the next
// dealer. Otherwise bidding's own finish() has already moved the game
// into the playing phase, so there's nothing left to do here.
func (m *Match) AfterBiddingResolved() (redealt bool) {
	if m.Bidding.BestCall == nil {
		m.Dealer = (m.Dealer + 1) % 4
		m.StartHand()
		return true
	}
	return false
}

// FinishHand scores the just-completed hand, folds the result into the
// match's running totals, and rotates the dealer for next time. Scores
// every detected sequence/carre/belot automatically - use
// FinishHandWithDeclared for the real declare-to-count behavior.
func (m *Match) FinishHand() HandResult {
	return m.finishHand(nil)
}

// FinishHandWithDeclared is the real production path: only credits a
// sequence/carre/belot the relevant team actually declared.
func (m *Match) FinishHandWithDeclared(declared DeclaredAnnounces) HandResult {
	return m.finishHand(&declared)
}

func (m *Match) finishHand(declared *DeclaredAnnounces) HandResult {
	var res HandResult
	if declared != nil {
		res = ScoreHandWithDeclared(m.Current, *declared)
	} else {
		res = ScoreHand(m.Current)
	}
	m.TeamScores[TeamA] = m.Current.Teams[TeamA].Score
	m.TeamScores[TeamB] = m.Current.Teams[TeamB].Score
	m.History = append(m.History, res)
	m.Dealer = (m.Dealer + 1) % 4
	return res
}

// MatchOver reports whether a team has reached the target score. Per
// standard rules the match continues until the END of the hand that
// crosses the threshold (already true here since we check after
// FinishHand), and if both teams cross in the same hand the higher
// score wins.
//
// This does NOT implement rule 9 ("с капо не се излиза" - you can't
// win the match directly off a валат/capo hand) by itself - use
// FinishHandAndCheckOver, which wraps FinishHand + this rule together
// so callers can't accidentally skip it.
func (m *Match) MatchOver() (over bool, winner TeamID) {
	aOver := m.TeamScores[TeamA] >= m.TargetScore
	bOver := m.TeamScores[TeamB] >= m.TargetScore
	if !aOver && !bOver {
		return false, 0
	}
	if m.TeamScores[TeamA] >= m.TeamScores[TeamB] {
		return true, TeamA
	}
	return true, TeamB
}

// FinishHandAndCheckOver scores the just-completed hand and reports
// whether the match is over - applying rule 9 along the way: a team
// can't clinch the match on a hand that included a валат/capo. If this
// hand would otherwise end the match AND it had a capo, the match is
// kept alive (over=false) so one more hand gets played; the same check
// re-applies after that hand too, so a chain of capo hands correctly
// keeps deferring the win until a "normal" hand decides it.
func (m *Match) FinishHandAndCheckOver() (res HandResult, over bool, winner TeamID) {
	res = m.FinishHand()
	over, winner = m.MatchOver()
	if over && res.CapoFound {
		return res, false, 0
	}
	return res, over, winner
}

// FinishHandAndCheckOverWithDeclared is FinishHandAndCheckOver's real
// production counterpart - only credits announces that were actually
// declared during play.
func (m *Match) FinishHandAndCheckOverWithDeclared(declared DeclaredAnnounces) (res HandResult, over bool, winner TeamID) {
	res = m.FinishHandWithDeclared(declared)
	over, winner = m.MatchOver()
	if over && res.CapoFound {
		return res, false, 0
	}
	return res, over, winner
}
