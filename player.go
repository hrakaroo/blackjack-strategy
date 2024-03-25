package main

type Action int

const (
	Stand Action = iota + 1
	Hit
	Split
	Double
)

type Result int

const (
	BlackJack = iota + 1
	Bust
	Lose
)

type Player struct {
	brain Brain

	hands      []*Hand
	handIndex  int
	doubleDown bool

	// Keep a running total of how much we have wagered
	wager int

	// Wins and losses
	bankroll int
}

func (p *Player) NewHand(newShoe bool) {
	bet := p.brain.Bet()

	p.bankroll -= bet
	p.hands = []*Hand{NewHand(bet)}
}

func (p *Player) Strategy() string {
	return p.brain.Name()
}

func (p *Player) Action(dealer Card) Action {

	if p.doubleDown {
		// todo - We already took the card so advance the handIndex

		p.doubleDown = false
		return Stand
	}

	action := p.brain.Action(dealer, p.hands[0])
	if action == Double {
		// Double the bet in the hand and only take one more card
		p.hands[0].bet *= 2
		p.doubleDown = true
		return Hit
	}
	return action
}

func (p *Player) Wager() int {
	return p.wager
}

func (p *Player) Bankroll() int {
	return p.bankroll
}

func (p *Player) Take(card Card, faceUp bool) {
	p.hands[p.handIndex].Take(card)
}

func (p *Player) DealerHasBlackJack() {
	if p.hands[0].IsBlackJack() {
		// Push
		p.bankroll += p.hands[0].bet
	}
}

func (p *Player) DealerHas(dealerTotal int) {
	for i := 0; i < len(p.hands); i++ {
		p.wager += p.hands[i].bet
		playerTotal, _ := p.hands[i].Total()
		if playerTotal > 21 {
			// The player busted, we lose what ever money we bet
		} else if p.hands[i].IsBlackJack() {
			// The player had blackjack dealer pays 2/3
			p.bankroll += p.hands[i].bet * 5 / 2
		} else if dealerTotal > 21 || playerTotal > dealerTotal {
			// Dealer bust, win 2x what we bet
			p.bankroll += p.hands[i].bet * 2
		}
	}
}

func NewPlayer(brain Brain) *Player {
	return &Player{brain: brain}
}
