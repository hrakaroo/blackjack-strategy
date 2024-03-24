package main

type Action int

const (
	Stand Action = iota + 1
	Hit
	Split
)

type Result int

const (
	BlackJack = iota + 1
	Bust
	Lose
)

type Player interface {
	Strategy() string

	// A new hand, newShoe indicates that we are dealing from a new shoe
	NewHand(newShoe bool)

	// Knowing the dealers face up card, what do they want to do, hit or stay
	Action(dealer Card) Action

	// The dealer give them the given card
	Take(card Card, faceUp bool)

	// Request to view the card facing up
	TopCard() Card

	DealerHasBlackJack()

	DealerHas(total int)

	Wager() int

	Bankroll() int
}
