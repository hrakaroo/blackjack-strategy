package brains

import "github.com/hrakaroo/blackjack-strategy/internal/game"

// Hit Soft 17 keeps hitting until we are at a hard 17 or better
type HitSoft17 struct {
}

func (b *HitSoft17) Name() string {
	return "Hit Soft 17"
}

func (b *HitSoft17) Bet() int {
	return 2
}

func (b *HitSoft17) Action(dealerCard game.Card, hand *game.Hand) game.Action {
	total, soft := hand.Total()
	if total < 17 || total == 17 && soft {
		return game.Hit
	}
	return game.Stand
}

func NewHitSoft17() game.Brain {
	return &HitSoft17{}
}
