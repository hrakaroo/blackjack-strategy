package main

import (
	"fmt"

	"github.com/hrakaroo/blackjack-strategy/internal/brains"
	"github.com/hrakaroo/blackjack-strategy/internal/game"
)

const DecksInShoe = 3

func main() {
	// Create our dealer
	dealer := game.NewDealer(brains.NewHitSoft17())

	// Create eyes for card countiung
	eyes := game.NewEyes()

	// Create the players
	players := []*game.Player{
		game.NewPlayer(brains.NewHitSoft17()),
		game.NewPlayer(brains.NewNeverBust()),
		game.NewPlayer(brains.NewSimple1()),
		game.NewPlayer(brains.NewSimple2()),
		game.NewPlayer(brains.NewPerfect()),
		game.NewPlayer(brains.NewFullCounting(eyes)),
	}

	var shoe *game.Shoe
	for i := 0; i < 1_000_000; i++ {
		var newShoe bool
		if shoe == nil || shoe.IsDone() {
			// fmt.Println("New shoe")
			shoe = game.NewShoe(DecksInShoe)
			shoe.Shuffle()
			newShoe = true
			eyes.NewShoe(DecksInShoe)
		}

		pullFn := shoe.Pull
		// Watch the shoe.Pull
		pullFn = eyes.Watch(pullFn)

		// Inform the players (and dealer) of a new hand
		for i := 0; i < len(players); i++ {
			players[i].NewHand(newShoe)
		}
		dealer.NewHand(newShoe)

		// Deal two cards to each
		for i := 0; i < 2; i++ {
			for p := 0; p < len(players); p++ {
				players[p].Deal(pullFn)
			}
			// Dealer
			dealer.Deal(pullFn)
		}

		// Check if the dealer has blackjack. No need to model insurance
		if dealer.Total() == 21 {
			for p := 0; p < len(players); p++ {
				players[p].DealerHasBlackJack()
			}
			continue
		}

		// Each Player goes
		for p := 0; p < len(players); p++ {
			players[p].Play(dealer.TopCard(), pullFn)
		}

		// Dealer goes
		dealer.Play(dealer.TopCard(), pullFn)
		dealerTotal := dealer.Total()

		// Update each player on the results of the dealer
		for p := 0; p < len(players); p++ {
			players[p].DealerHas(dealerTotal)
		}
	}

	for i := 0; i < len(players); i++ {
		player := players[i]
		wagers := player.Wagers()
		wins := player.Wins()
		ratio := float32(wins) * 100.0 / float32(wagers)
		fmt.Printf("Player(%d) - %-20s - %d : %d = %0.2f%%\n", i, player.Strategy(), wagers, wins, ratio)
	}
}
