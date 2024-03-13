package main

import "fmt"

type Action int

const (
	Stay Action = iota + 1
	Hit
	Bust
)

type Player interface {
	// A new hand
	NewHand()
	Total() int
	Action(dealer Card) Action
	Take(card Card)
	// Request to view the card facing up
	TopCard() Card
	// A new shoe is triggered
	NewShoe()
	String() string
}

type HitSoft17 struct {
	hand *Hand
}

func (d *HitSoft17) NewHand() {
	d.hand = &Hand{}
}

func (d *HitSoft17) Total() int {
	// Keep hitting until we get to hard 17 or better
	total, _ := d.hand.Total()
	return total
}

func (d *HitSoft17) Action(dealer Card) Action {
	// Keep hitting until we get to hard 17 or better
	total, soft := d.hand.Total()
	if total > 21 {
		return Bust
	}
	if total < 17 || total == 17 && soft {
		return Hit
	}
	return Stay
}

func (d *HitSoft17) Take(card Card) {
	d.hand.Take(card)
}

func (d *HitSoft17) TopCard() Card {
	return d.hand.Cards[0]
}

func (d *HitSoft17) NewShoe() {
	// We don't care
}

func (d *HitSoft17) String() string {
	return fmt.Sprintf("%v", *d.hand)
}

func NewHitSoft17() Player {
	return &HitSoft17{}
}
