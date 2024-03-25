package main

type Dealer struct {
	Player
}

func (d *Dealer) Total() int {
	total, _ := d.hands[0].Total()
	return total
}

func (d *Dealer) TopCard() Card {
	return d.hands[0].Cards[0]
}

func NewDealer() *Dealer {
	d := &Dealer{}
	d.brain = NewHitSoft17()
	return d
}
