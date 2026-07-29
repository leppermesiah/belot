package engine

import "fmt"

// Suit represents one of the four card suits.
type Suit int

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

func (s Suit) String() string {
	switch s {
	case Clubs:
		return "♣"
	case Diamonds:
		return "♦"
	case Hearts:
		return "♥"
	case Spades:
		return "♠"
	}
	return "?"
}

var AllSuits = []Suit{Clubs, Diamonds, Hearts, Spades}

// Rank represents a card rank. Belot uses 32 cards: 7,8,9,10,J,Q,K,A.
type Rank int

const (
	Seven Rank = iota
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

func (r Rank) String() string {
	switch r {
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "10"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Ace:
		return "A"
	}
	return "?"
}

var AllRanks = []Rank{Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

// Card values differ depending on whether its suit is trump or not.
// Non-trump order (low->high): 7,8,9,J,Q,K,10,A
// Trump order    (low->high): 7,8,Q,K,10,A,9,J

// NonTrumpValue returns card point value when its suit is NOT trump.
func (r Rank) NonTrumpValue() int {
	switch r {
	case Ace:
		return 11
	case Ten:
		return 10
	case King:
		return 4
	case Queen:
		return 3
	case Jack:
		return 2
	default: // 7, 8, 9
		return 0
	}
}

// TrumpValue returns card point value when its suit IS trump.
func (r Rank) TrumpValue() int {
	switch r {
	case Jack:
		return 20
	case Nine:
		return 14
	case Ace:
		return 11
	case Ten:
		return 10
	case King:
		return 4
	case Queen:
		return 3
	default: // 7, 8
		return 0
	}
}

// nonTrumpRankOrder maps rank -> strength when suit is not trump (higher wins).
var nonTrumpRankOrder = map[Rank]int{
	Seven: 0, Eight: 1, Nine: 2, Jack: 3, Queen: 4, King: 5, Ten: 6, Ace: 7,
}

// trumpRankOrder maps rank -> strength when suit is trump (higher wins).
var trumpRankOrder = map[Rank]int{
	Seven: 0, Eight: 1, Queen: 2, King: 3, Ten: 4, Ace: 5, Nine: 6, Jack: 7,
}

// Card is a single playing card.
type Card struct {
	Suit Suit
	Rank Rank
}

func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.Rank, c.Suit)
}

// Value returns the point value of this card given the current trump suit.
func (c Card) Value(trump Suit) int {
	if c.Suit == trump {
		return c.Rank.TrumpValue()
	}
	return c.Rank.NonTrumpValue()
}

// StrengthAgainst returns this card's relative strength for trick-taking
// purposes, given the current trump suit. Only meaningful when comparing
// cards of the same suit (or both trump).
func (c Card) Strength(trump Suit) int {
	if c.Suit == trump {
		return trumpRankOrder[c.Rank]
	}
	return nonTrumpRankOrder[c.Rank]
}

// NewDeck returns a full 32-card Belot deck, unshuffled.
func NewDeck() []Card {
	deck := make([]Card, 0, 32)
	for _, s := range AllSuits {
		for _, r := range AllRanks {
			deck = append(deck, Card{Suit: s, Rank: r})
		}
	}
	return deck
}
