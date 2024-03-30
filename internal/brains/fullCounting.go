package brains

import (
	"fmt"

	"github.com/hrakaroo/blackjack-strategy/internal/game"
)

// FullCounting follows the perfect strategy and does perfect card counting.
type FullCounting struct {
	runningCount int
	dealCount    int
	decksInShoe  int
}

func (b *FullCounting) Name() string {
	return "Full Counting"
}

func (b *FullCounting) Bet() int {
	bet := 2

	// Compute real count
	deckRemaining := ((b.decksInShoe*52 - b.dealCount) / 52) + 1

	realCount := b.runningCount / deckRemaining

	if realCount > 15 {
		fmt.Println("Big bet")
		return 10 * bet
	}
	return bet
}

func (b *FullCounting) Take(card game.Card) {
	b.dealCount++
	if card.Rank >= 2 && card.Rank <= 6 {
		b.runningCount++
	} else if card.Rank == 1 || card.Rank >= 10 {
		b.runningCount--
	}
}

func (b *FullCounting) NewShoe(decksInShoe int) {
	b.runningCount = 0
	b.dealCount = 0
	b.decksInShoe = decksInShoe
}

func (b *FullCounting) Action(dealerCard game.Card, hand *game.Hand) game.Action {
	total, soft := hand.Total()

	if total > 21 {
		// Can't keep hitting if we have busted
		return game.Stand
	}

	firstTwo := len(hand.Cards) == 2

	// Check for pairs
	if firstTwo && hand.Cards[0].Rank == hand.Cards[1].Rank {
		card := hand.Cards[0].Rank
		if card == 9 && ((dealerCard.Rank >= 2 && dealerCard.Rank <= 6) || dealerCard.Rank == 8 || dealerCard.Rank == 9) {
			return game.Split
		}
		if card == 8 {
			return game.Split
		}
		if card == 7 && dealerCard.Rank >= 2 && dealerCard.Rank <= 7 {
			return game.Split
		}
		if card == 6 && dealerCard.Rank >= 3 && dealerCard.Rank <= 6 {
			return game.Split
		}
		if card == 3 && dealerCard.Rank >= 4 && dealerCard.Rank <= 7 {
			return game.Split
		}
		if card == 2 && dealerCard.Rank >= 4 && dealerCard.Rank <= 7 {
			return game.Split
		}
		if card == 1 {
			// Split the Aces
			return game.Split
		}

	}

	if soft {
		switch total {
		case 21:
			fallthrough
		case 20:
			return game.Stand
		case 19:
			if firstTwo && dealerCard.Rank == 6 {
				return game.Double
			}
			return game.Stand
		case 18:
			if firstTwo && dealerCard.Rank <= 6 {
				return game.Double
			}
			if dealerCard.Rank <= 8 {
				return game.Stand
			}
		case 17:
			if firstTwo && dealerCard.Rank >= 3 && dealerCard.Rank <= 6 {
				return game.Double
			}
		case 16:
			fallthrough
		case 15:
			if firstTwo && dealerCard.Rank >= 4 && dealerCard.Rank <= 6 {
				return game.Double
			}
		case 14:
			fallthrough
		case 13:
			if firstTwo && dealerCard.Rank >= 5 && dealerCard.Rank <= 6 {
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
			if firstTwo {
				return game.Double
			}
		case 10:
			if firstTwo && dealerCard.Rank <= 9 {
				return game.Double
			}
		case 9:
			if firstTwo && dealerCard.Rank >= 3 && dealerCard.Rank <= 6 {
				return game.Double
			}
		}
	}
	return game.Hit
}

func NewFullCounting(eyes *game.Eyes) game.Brain {
	brain := &FullCounting{}
	// Register call backs
	eyes.OnTake(brain.Take)
	eyes.OnNewShoe(brain.NewShoe)

	return brain
}
