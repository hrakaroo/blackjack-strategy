package game

import "fmt"

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

type Brain interface {
	Name() string
	Bet() int
	Action(dealerCard Card, hand *Hand) Action
}

type Player struct {
	brain Brain

	hands []*Hand

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

func (p *Player) Deal(take func() Card) {
	// Take a card
	p.hands[0].Take(take())
}

func (p *Player) Play(dealer Card, take func() Card) {
	// Go through each hand
	for i := 0; i < len(p.hands); i++ {
		// Dereference the current hand we are working on
		hand := p.hands[i]

		// Keep taking cards until we are done
	PROMPT:
		for {
			if len(hand.Cards) == 1 {
				// This is due to a split
				hand.Take(take())
				continue PROMPT
			}

			// Determine what the course of action is for this hand
			action := p.brain.Action(dealer, hand)
			switch action {
			case Stand:
				break PROMPT
			case Double:
				// Belt and suspenders, make sure we can double and this is not a bug
				if len(hand.Cards) != 2 {
					panic("Illegal double")
				}
				// Double our bet and only take one more hand
				hand.bet *= 2
				hand.Take(take())
				break PROMPT
			case Hit:
				hand.Take(take())
			case Split:
				// Belt and suspenders, make sure we can split and this is not a bug
				if len(hand.Cards) != 2 {
					panic("Illegal split")
				}
				card1 := hand.Cards[0]
				card2 := hand.Cards[1]

				// Keep card 1 in the current hand
				hand.Cards = []Card{card1}

				// Put card 2 in the new hand
				newHand := NewHand(p.hands[i].bet)
				newHand.Cards = []Card{card2}

				// Put the new hand at the end.  Technically this isn't how it works at a casino but
				//  as this hand would then get played last and not next but it shouldn't change the odds
				p.hands = append(p.hands, newHand)
			default:
				panic(fmt.Sprintf("Forgot to handle %v", action))
			}
		}
	}
}

func (p *Player) Wagers() int {
	return p.wagers
}

func (p *Player) Wins() int {
	return p.wins
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
