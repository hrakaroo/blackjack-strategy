package brains

import "github.com/hrakaroo/blackjack-strategy/internal/game"

type Simple2 struct {
}

func (b *Simple2) Name() string {
	return "No Split"
}

func (b *Simple2) Bet() int {
	return 2
}

func (b *Simple2) Action(dealerCard game.Card, hand *game.Hand) game.Action {
	total, soft := hand.Total()

	if total > 21 {
		return game.Stand
	}

	canDouble := len(hand.Cards) == 2

	if soft {
		switch total {
		case 21:
			fallthrough
		case 20:
			return game.Stand
		case 19:
			if canDouble && dealerCard.Rank == 6 {
				return game.Double
			}
			return game.Stand
		case 18:
			if canDouble && dealerCard.Rank <= 6 {
				return game.Double
			}
			if dealerCard.Rank <= 8 {
				return game.Stand
			}
		case 17:
			if canDouble && dealerCard.Rank >= 3 && dealerCard.Rank <= 6 {
				return game.Double
			}
		case 16:
			fallthrough
		case 15:
			if canDouble && dealerCard.Rank >= 4 && dealerCard.Rank <= 6 {
				return game.Double
			}
		case 14:
			fallthrough
		case 13:
			if canDouble && dealerCard.Rank >= 5 && dealerCard.Rank <= 6 {
				return game.Double
			}
		}
	} else {
		switch total {
		case 21:
			fallthrough
		case 20:
			fallthrough
		case 19:
			fallthrough
		case 18:
			fallthrough
		case 17:
			return game.Stand
		case 16:
			fallthrough
		case 15:
			fallthrough
		case 14:
			fallthrough
		case 13:
			if dealerCard.Rank <= 6 {
				return game.Stand
			}
		case 12:
			if dealerCard.Rank >= 4 && dealerCard.Rank <= 6 {
				return game.Stand
			}
		case 11:
			if canDouble {
				return game.Double
			}
		case 10:
			if canDouble && dealerCard.Rank <= 9 {
				return game.Double
			}
		case 9:
			if canDouble && dealerCard.Rank >= 3 && dealerCard.Rank <= 6 {
				return game.Double
			}
		}
	}
	return game.Hit
}

func NewSimple2() game.Brain {
	return &Simple2{}
}
