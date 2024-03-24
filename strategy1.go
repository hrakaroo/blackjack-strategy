package main

type Strategy1 struct {
	hand *Hand

	BasicPlayer
}

func (d *Strategy1) Strategy() string {
	return "Strategy 1"
}

func (d *Strategy1) NewHand(newShoe bool) {
	// Bet $2
	bet := 2

	d.wager += bet
	d.bankroll -= bet
	d.hand = NewHand(bet)
}

func (d *Strategy1) Total() int {
	total, _ := d.hand.Total()
	return total
}

func (d *Strategy1) Action(dealer Card) Action {

	total, _ := d.hand.Total()
	if total >= 17 {
		// Always stand on 18 or better
		return Stand
	}

	if total >= 13 {
		if dealer.Rank > 7 {
			return Hit
		}
		return Stand
	}

	if total == 12 {
		if dealer.Rank >= 4 && dealer.Rank <= 6 {
			return Stand
		}
		return Hit
	}

	return Hit
}

func (d *Strategy1) Take(card Card, faceUp bool) {
	d.hand.Take(card)
}

func (d *Strategy1) TopCard() Card {
	return d.hand.Cards[0]
}

func (d *Strategy1) DealerHasBlackJack() {
	if d.hand.IsBlackJack() {
		// Push
		d.bankroll += d.hand.bet
	}
}

func (d *Strategy1) DealerHas(dealerTotal int) {
	playerTotal, _ := d.hand.Total()
	if playerTotal > 21 {
		// The player busted, we lose what ever money we bet
	} else if d.hand.IsBlackJack() {
		// The player had blackjack dealer pays 2/3
		d.bankroll += d.hand.bet * 5 / 2
	} else if dealerTotal > 21 || playerTotal > dealerTotal {
		// Dealer bust, win 2x what we bet
		d.bankroll += d.hand.bet * 2
	}
}

func NewStrategy1() Player {
	return &Strategy1{}
}
