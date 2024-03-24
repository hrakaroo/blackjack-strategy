package main

type HitSoft17 struct {
	hand *Hand

	BasicPlayer
}

func (d *HitSoft17) Strategy() string {
	return "Hit Soft 17"
}

func (d *HitSoft17) NewHand(newShoe bool) {
	// Bet $2
	bet := 2

	d.wager += bet
	d.bankroll -= bet
	d.hand = NewHand(bet)
}

func (d *HitSoft17) Action(dealer Card) Action {
	// Keep hitting until we get to hard 17 or better
	total, soft := d.hand.Total()
	if total < 17 || total == 17 && soft {
		return Hit
	}
	return Stand
}

func (d *HitSoft17) Take(card Card, faceUp bool) {
	d.hand.Take(card)
}

func (d *HitSoft17) TopCard() Card {
	return d.hand.Cards[0]
}

func (d *HitSoft17) DealerHasBlackJack() {
	if d.hand.IsBlackJack() {
		// Push
		d.bankroll += d.hand.bet
	}
}

func (d *HitSoft17) DealerHas(dealerTotal int) {
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

func NewHitSoft17() Player {
	return &HitSoft17{}
}
