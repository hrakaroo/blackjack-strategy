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
	wagers int

	// Our total for wins
	wins int
}

func (p *Player) NewHand(newShoe bool) {
	bet := p.brain.Bet()

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
		// Belt and suspenders, make sure we can double and this is not a bug
		if len(p.hands[0].Cards) != 2 {
			panic("Illegal double")
		}

		// Double the bet in the hand and only take one more card
		p.hands[0].bet *= 2
		p.doubleDown = true
		return Hit
	}
	return action
}

func (p *Player) Wagers() int {
	return p.wagers
}

func (p *Player) Wins() int {
	return p.wins
}

func (p *Player) Take(card Card, faceUp bool) {
	p.hands[p.handIndex].Take(card)
}

func (p *Player) DealerHasBlackJack() {
	// Technically the loop is unnecessary as we can't split if the dealer has blackjack but
	//  for completeness ...
	for i := 0; i < len(p.hands); i++ {
		p.wagers += p.hands[i].bet
		if p.hands[0].IsBlackJack() {
			// Push
			p.wins += p.hands[i].bet
		}
	}
}

func (p *Player) DealerHas(dealerTotal int) {
	for i := 0; i < len(p.hands); i++ {
		p.wagers += p.hands[i].bet
		playerTotal, _ := p.hands[i].Total()

		if playerTotal > 21 {
			// The player busted, they lose what ever money they bet
			continue
		}

		if p.hands[i].IsBlackJack() {
			// The player had blackjack, dealer pays 2/3
			//  So a bet of $2 would pay $3 for a total reclaim of $5
			//  $2 * 5 = $10 / 2 = $5
			p.wins += p.hands[i].bet * 5 / 2
			continue
		}

		if dealerTotal > 21 || playerTotal > dealerTotal {
			// Dealer bust, win 2x what we bet
			p.wins += p.hands[i].bet * 2
			continue
		}

		if dealerTotal == playerTotal {
			// Push, just get our bet back
			p.wins += p.hands[i].bet
			continue
		}
	}
}

func NewPlayer(brain Brain) *Player {
	return &Player{brain: brain}
}
