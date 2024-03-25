package game

import "fmt"

type Rank int

func (r Rank) String() string {
	switch int(r) {
	case 1:
		return "A"
	case 10:
		return "T"
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	default:
		return fmt.Sprintf("%d", int(r))
	}
}

func (r Rank) Value() []int {
	if r == 1 {
		return []int{1, 11}
	}
	if r > 10 {
		return []int{10}
	}
	return []int{int(r)}
}
