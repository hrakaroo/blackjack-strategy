package main

type Brain interface {
	Name() string
	Bet() int
	Action(dealerCard Card, hand *Hand) Action
}

type HitSoft17 struct {
}

func (b *HitSoft17) Name() string {
	return "Hit Soft 17"
}

func (b *HitSoft17) Bet() int {
	return 2
}

func (b *HitSoft17) Action(dealerCard Card, hand *Hand) Action {
	total, soft := hand.Total()
	if total < 17 || total == 17 && soft {
		return Hit
	}
	return Stand
}

func NewHitSoft17() Brain {
	return &HitSoft17{}
}

type Simple struct {
}

func (b *Simple) Name() string {
	return "Simple"
}

func (b *Simple) Bet() int {
	return 2
}

func (b *Simple) Action(dealerCard Card, hand *Hand) Action {
	total, soft := hand.Total()
	if soft {
		if total >= 19 {
			return Stand
		}

		if total == 18 {
			if dealerCard.Rank <= 8 {
				return Stand
			} else {
				return Hit
			}
		}
	} else {
		if total >= 17 {
			// Always stand on 18 or better
			return Stand
		}

		if total >= 13 {
			if dealerCard.Rank > 7 {
				return Hit
			}
			return Stand
		}

		if total == 12 {
			if dealerCard.Rank >= 4 && dealerCard.Rank <= 6 {
				return Stand
			}
			return Hit
		}
	}

	return Hit
}

func NewSimple() Brain {
	return &Simple{}
}
