package main

type Suit int

func (s Suit) String() string {
	return [...]string{"H", "C", "D", "S"}[s-1]
}

func (s Suit) EnumIndex() int {
	return int(s)
}
