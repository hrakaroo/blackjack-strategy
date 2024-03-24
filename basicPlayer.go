package main

type BasicPlayer struct {
	// Keep a running total of how much we have wagered
	wager int

	// Wins and losses
	bankroll int
}

func (d *BasicPlayer) Wager() int {
	return d.wager
}

func (d *BasicPlayer) Bankroll() int {
	return d.bankroll
}
