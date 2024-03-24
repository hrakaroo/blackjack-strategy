package main

type Dealer struct {
	HitSoft17
}

func (d *Dealer) Total() int {
	total, _ := d.hand.Total()
	return total
}

func NewDealer() *Dealer {
	return &Dealer{}
}
