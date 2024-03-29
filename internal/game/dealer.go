package game

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

func NewDealer(brain Brain) *Dealer {
	d := &Dealer{}
	d.brain = brain
	return d
}
