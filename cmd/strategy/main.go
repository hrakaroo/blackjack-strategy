package main

import (
	"fmt"

	"github.com/hrakaroo/blackjack-strategy/internal/brains"
	"github.com/hrakaroo/blackjack-strategy/internal/game"
)

func main() {
	// Create our dealer
	dealer := game.NewDealer(brains.NewHitSoft17())

	// Create the players
	players := []*game.Player{
		game.NewPlayer(brains.NewHitSoft17()),
		game.NewPlayer(brains.NewSimple1()),
		game.NewPlayer(brains.NewSimple2()),
		game.NewPlayer(brains.NewPerfect()),
		game.NewPlayer(brains.NewPerfect()),
	}

	var shoe *game.Shoe
	for i := 0; i < 1_000_000; i++ {
		var newShoe bool
		if shoe == nil || shoe.IsDone() {
			// fmt.Println("New shoe")
			shoe = game.NewShoe(3)
			shoe.Shuffle()
			newShoe = true
		}

		// Inform the players (and dealer) of a new hand
		for i := 0; i < len(players); i++ {
			players[i].NewHand(newShoe)
		}
		dealer.NewHand(newShoe)

		// Deal two cards to each
		for i := 0; i < 2; i++ {
			for p := 0; p < len(players); p++ {
				players[p].Take(shoe.Pull(), true)
			}
			// Dealer
			dealer.Take(shoe.Pull(), i == 0)
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
			player := players[p]
			for {
				action := player.Action(dealer.TopCard())
				if action == game.Hit {
					player.Take(shoe.Pull(), true)
				} else {
					break
				}
			}
		}

		// Dealer goes
		for {
			action := dealer.Action(dealer.TopCard())
			if action == game.Hit {
				dealer.Take(shoe.Pull(), true)
			} else {
				break
			}
		}
		dealerTotal := dealer.Total()

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
