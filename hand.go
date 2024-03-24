package main

import (
	"cmp"
	"slices"
)

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
	bet   int
}

func NewHand(bet int) *Hand {
	return &Hand{bet: bet}
}

func (h *Hand) Take(c Card) {
	// Receive the card
	h.Cards = append(h.Cards, c)
}

func (h *Hand) IsBlackJack() bool {
	total, _ := h.Total()
	return len(h.Cards) == 2 && total == 21
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
