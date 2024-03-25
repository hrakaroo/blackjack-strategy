package main

type Brain interface {
	Name() string
	Bet() int
	Action(dealerCard Card, hand *Hand) Action
}

// Hit Soft 17 keeps hitting until we are at a hard 17 or better
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

type Simple1 struct {
}

func (b *Simple1) Name() string {
	return "Simple 1"
}

func (b *Simple1) Bet() int {
	return 2
}

func (b *Simple1) Action(dealerCard Card, hand *Hand) Action {
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
		// Hard
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

func NewSimple1() Brain {
	return &Simple1{}
}

type Simple2 struct {
}

func (b *Simple2) Name() string {
	return "Simple 2"
}

func (b *Simple2) Bet() int {
	return 2
}

func (b *Simple2) Action(dealerCard Card, hand *Hand) Action {
	total, soft := hand.Total()

	if total > 21 {
		return Stand
	}

	canDouble := len(hand.Cards) == 2

	if soft {
		switch total {
		case 21:
			fallthrough
		case 20:
			return Stand
		case 19:
			if canDouble && dealerCard.Rank == 6 {
				return Double
			}
			return Stand
		case 18:
			if canDouble && dealerCard.Rank <= 6 {
				return Double
			}
			if dealerCard.Rank <= 8 {
				return Stand
			}
		case 17:
			if canDouble && dealerCard.Rank >= 3 && dealerCard.Rank <= 6 {
				return Double
			}
		case 16:
			fallthrough
		case 15:
			if canDouble && dealerCard.Rank >= 4 && dealerCard.Rank <= 6 {
				return Double
			}
		case 14:
			fallthrough
		case 13:
			if canDouble && dealerCard.Rank >= 5 && dealerCard.Rank <= 6 {
				return Double
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
			return Stand
		case 16:
			fallthrough
		case 15:
			fallthrough
		case 14:
			fallthrough
		case 13:
			if dealerCard.Rank <= 6 {
				return Stand
			}
		case 12:
			if dealerCard.Rank >= 4 && dealerCard.Rank <= 6 {
				return Stand
			}
		case 11:
			if canDouble {
				return Double
			}
		case 10:
			if canDouble && dealerCard.Rank <= 9 {
				return Double
			}
		case 9:
			if canDouble && dealerCard.Rank >= 3 && dealerCard.Rank <= 6 {
				return Double
			}
		}
	}
	return Hit
}

func NewSimple2() Brain {
	return &Simple2{}
}
