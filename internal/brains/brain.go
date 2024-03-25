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
