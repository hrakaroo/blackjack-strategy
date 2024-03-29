package game

type Eyes struct {
	takeFn func() Card
}

func (e *Eyes) Take() Card {
	return e.takeFn()
}

func (e *Eyes) Watch(take func() Card) func() Card {
	e.takeFn = take
	return e.Take
}

func NewEyes() *Eyes {
	return &Eyes{}
}
