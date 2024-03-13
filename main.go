package main

import (
	"cmp"
	"fmt"
	"math/rand"
	"slices"
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

func (h *Hand) Total() (total int, soft bool) {
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
			return totals[i].Value, totals[i].Soft
		}
	}
	return totals[0].Value, totals[0].Soft
}

func main() {
	shoe := NewShoe(3)
	shoe.Shuffle()

	// Create our dealer
	dealer := NewHitSoft17()

	// Create the players
	players := []Player{NewHitSoft17(), NewHitSoft17(), NewHitSoft17()}
	scores := make([]int, len(players))
	count := 0

	for i := 0; i < 1_000_000; i++ {
		if shoe.IsDone() {
			// fmt.Println("New shoe")
			shoe = NewShoe(3)
			shoe.Shuffle()
		}

		for i := 0; i < len(players); i++ {
			players[i].NewHand()
		}
		dealer.NewHand()

		// Deal two cards to each
		for i := 0; i < 2; i++ {
			for p := 0; p < len(players); p++ {
				players[p].Take(shoe.Pull())
			}
			// Dealer
			dealer.Take(shoe.Pull())
		}

		// Each Player goes
		for p := 0; p < len(players); p++ {
			player := players[p]
			for {
				action := player.Action(dealer.TopCard())
				if action == Hit {
					player.Take(shoe.Pull())
				} else {
					break
				}
			}
		}
		// Dealer goes
		for {
			action := dealer.Action(dealer.TopCard())
			if action == Hit {
				dealer.Take(shoe.Pull())
			} else {
				break
			}
		}
		dealerTotal := dealer.Total()

		// fmt.Printf("Dealer = %v = %d\n", dealer, dealerTotal)

		// Determine winners
		for i := 0; i < len(players); i++ {
			total := players[i].Total()
			// fmt.Printf("Player(%d) = %v = %d\n", i, players[i], total)
			if total > 21 {
				// fmt.Printf("Player BUSTS\n")
				// A bust is always a loss
				scores[i]--
			} else if dealerTotal > 21 {
				// fmt.Printf("Dealer BUSTS\n")
				// Dealer bust is a win
				scores[i]++
			} else if total > dealerTotal {
				// fmt.Printf("Player WINS %d > %d \n", total, dealerTotal)
				scores[i]++
			} else if total < dealerTotal {
				// fmt.Printf("Dealer WINS %d < %d \n", total, dealerTotal)
				scores[i]--
			} else {
				// fmt.Printf("Push %d = %d\n", total, dealerTotal)
			}
		}
		count++
	}

	for i := 0; i < len(scores); i++ {
		fmt.Printf("Player %d/%d = %d%%\n", scores[i], count, (scores[i]+count)*100/count)
	}
}
