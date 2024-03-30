package game

type Eyes struct {
	takeFn func() Card

	takeCallback    []func(card Card)
	newShoeCallback []func(decksInShoe int)
}

func (e *Eyes) Take() Card {
	// Call the real function
	card := e.takeFn()

	// Send the card to the callbacks
	for i := 0; i < len(e.takeCallback); i++ {
		e.takeCallback[i](card)
	}

	// Return the card
	return card
}

func (e *Eyes) NewShoe(decksInShoe int) {
	// Notify all callbacks
	for i := 0; i < len(e.newShoeCallback); i++ {
		e.newShoeCallback[i](decksInShoe)
	}
}

func (e *Eyes) Watch(take func() Card) func() Card {
	// Save off the "real" function to call
	e.takeFn = take
	// Give them back our intercepted call
	return e.Take
}

func (e *Eyes) OnTake(callback func(card Card)) {
	e.takeCallback = append(e.takeCallback, callback)
}

func (e *Eyes) OnNewShoe(callback func(decksInShoe int)) {
	e.newShoeCallback = append(e.newShoeCallback, callback)
}

func NewEyes() *Eyes {
	return &Eyes{}
}
