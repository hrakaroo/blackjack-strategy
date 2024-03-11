package main

import (
	"cmp"
	"fmt"
	"math/rand"
	"slices"
	"strings"
	"time"
)

type Suit int

const (
	Heart Suit = iota + 1
	Club
	Diamond
	Spade
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (s Suit) String() string {
	return [...]string{"H", "C", "D", "S"}[s-1]
}

func (s Suit) EnumIndex() int {
	return int(s)
}

type Rank int

func (r Rank) String() string {
	switch int(r) {
	case 1:
		return "A"
	case 10:
		return "T"
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	default:
		return fmt.Sprintf("%d", int(r))
	}
}

func (r Rank) Value() []int {
	if r == 1 {
		return []int{1, 11}
	}
	if r > 10 {
		return []int{10}
	}
	return []int{int(r)}
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
func (s *Shoe) Pull() Card {
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

	// The end card is set at around the last 15%
	end := int(float32(len(cards)) * .85)
	return &Shoe{cards: cards, end: end}
}

type Total struct {
	Value int
	Soft  bool
}

func contains(s []Total, v Total) bool {
	for i := 0; i < len(s); i++ {
		if s[i].Value == v.Value {
			return true
		}
	}
	return false
}

type Hand struct {
	Cards []Card
}

func (h *Hand) Take(c Card) {
	// Receive the card
	h.Cards = append(h.Cards, c)
}

func (h *Hand) Total() (total int, soft bool, bust bool) {
	// Calculate our totals

	totals := []Total{{0, false}}
	for i := 0; i < len(h.Cards); i++ {
		values := h.Cards[i].Value()
		var newTotals []Total
		for t := 0; t < len(totals); t++ {
			t1 := Total{totals[t].Value + values[0], totals[t].Soft}
			if !contains(newTotals, t1) {
				newTotals = append(newTotals, t1)
			}
			if len(values) > 1 {
				t2 := Total{totals[t].Value + values[1], true}
				if !contains(newTotals, t2) {
					newTotals = append(newTotals, t2)
				}
			}
		}
		slices.SortFunc(newTotals, func(a, b Total) int { return cmp.Compare(a.Value, b.Value) })
		totals = newTotals
	}

	// Find the highest score not over 21
	for i := len(totals) - 1; i >= 0; i-- {
		if totals[i].Value <= 21 {
			return totals[i].Value, totals[i].Soft, false
		}
	}
	return totals[0].Value, totals[0].Soft, true
}

func main() {
	shoe := NewShoe(3)
	shoe.Shuffle()

	players := []*Hand{{}, {}}

	// Deal two cards to each
	for i := 0; i < 2; i++ {
		for p := 0; p < len(players); p++ {
			players[p].Take(shoe.Pull())
		}
	}

	// Each Player goes
	for p := 0; p < len(players)-1; p++ {
		player := players[p]
		for {
			// Keep hitting until we get to hard 17 or better
			total, soft, bust := player.Total()
			if bust {
				fmt.Printf("Player %d bust!\n", p)
				break
			}
			if total > 17 || total == 17 && !soft {
				break
			}
			player.Take(shoe.Pull())
		}
	}

	// Now the dealer
	dealer := players[len(players)-1]
	for {
		// Keep hitting until we get to hard 17 or better
		total, soft, bust := dealer.Total()
		if bust {
			fmt.Printf("Dealer bust: %d!\n", total)
			break
		}
		if total > 17 || total == 17 && !soft {
			break
		}
		dealer.Take(shoe.Pull())
	}

	for i := 0; i < len(players); i++ {
		total, _, _ := players[i].Total()
		fmt.Printf("%v = %d\n", *players[i], total)
	}
}
