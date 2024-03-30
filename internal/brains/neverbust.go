package brains

import "github.com/hrakaroo/blackjack-strategy/internal/game"

// Never bust is banking on the dealer busting and so tries to get the highest hand
//
//	possible without busting
type NeverBust struct {
}

func (b *NeverBust) Name() string {
	return "Never Bust"
}

func (b *NeverBust) Bet() int {
	return 2
}

func (b *NeverBust) Action(dealerCard game.Card, hand *game.Hand) game.Action {
	total, soft := hand.Total()
	if (total == 10 || total == 11) && len(hand.Cards) == 2 {
		return game.Double
	}
	if total <= 11 {
		return game.Hit
	}
	if total >= 17 {
		return game.Stand
	}
	if soft {
		return game.Hit
	}
	return game.Stand
}

func NewNeverBust() game.Brain {
	return &NeverBust{}
}
