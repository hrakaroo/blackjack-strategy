package brains

import "github.com/hrakaroo/blackjack-strategy/internal/game"

// Simple 1 follows the strategoy at
//  https://www.blackjackapprenticeship.com/blackjack-strategy-charts/
//  but doesn't split or double
type Simple1 struct {
}

func (b *Simple1) Name() string {
	return "No Split/No Double"
}

func (b *Simple1) Bet() int {
	return 2
}

func (b *Simple1) Action(dealerCard game.Card, hand *game.Hand) game.Action {
	total, soft := hand.Total()
	if soft {
		if total >= 19 {
			return game.Stand
		}

		if total == 18 {
			if dealerCard.Rank <= 8 {
				return game.Stand
			} else {
				return game.Hit
			}
		}
	} else {
		// Hard
		if total >= 17 {
			// Always stand on 18 or better
			return game.Stand
		}

		if total >= 13 {
			if dealerCard.Rank > 7 {
				return game.Hit
			}
			return game.Stand
		}

		if total == 12 {
			if dealerCard.Rank >= 4 && dealerCard.Rank <= 6 {
				return game.Stand
			}
			return game.Hit
		}
	}

	return game.Hit
}

func NewSimple1() game.Brain {
	return &Simple1{}
}