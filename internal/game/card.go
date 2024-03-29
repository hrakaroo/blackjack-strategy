package game

import (
	"math/rand"
	"strings"
	"time"
)

const (
	Heart Suit = iota + 1
	Club
	Diamond
	Spade
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Card struct {
	Suit Suit
	Rank Rank
}

func (c Card) String() string {
	return c.Rank.String() + " " + c.Suit.String()
}

func (c Card) Value() []int {
	return c.Rank.Value()
}

type Deck [52]Card

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
}

func (d *Deck) String() string {
	strs := make([]string, len(d))
	for i := 0; i < len(d); i++ {
		strs[i] = d[i].String()
	}
	return strings.Join(strs, ", ")
}

func NewDeck() *Deck {
	d := Deck{}
	index := 0
	for s := Heart; s <= Spade; s++ {
		for v := 1; v <= 13; v++ {
			d[index] = Card{Suit: s, Rank: Rank(v)}
			index++
		}
	}
	return &d
}

type Shoe struct {
	index int
	end   int
	cards []Card
}

func (s *Shoe) Shuffle() {
	rand.Shuffle(len(s.cards), func(i, j int) { s.cards[i], s.cards[j] = s.cards[j], s.cards[i] })
}

func (s *Shoe) String() string {
	strs := make([]string, len(s.cards))
	for i := 0; i < len(s.cards); i++ {
		strs[i] = s.cards[i].String()
	}
	return strings.Join(strs, ", ")
}

// Pull a card out of the Shoe
func (s *Shoe) Pull(faceDown bool) Card {
	if s.index > len(s.cards) {
		panic("Went past deck")
	}
	c := s.cards[s.index]
	s.index++
	return c
}

func (s *Shoe) IsDone() bool {
	return s.index > s.end
}

func NewShoe(decks int) *Shoe {
	cards := make([]Card, decks*52)
	index := 0

	for i := 0; i < decks; i++ {
		d := NewDeck()
		for i := 0; i < len(d); i++ {
			cards[index] = d[i]
			index++
		}
	}

	// The end card is set at around the last 20%
	end := int(float32(len(cards)) * .80)
	return &Shoe{cards: cards, end: end}
}
