package main

import (
	"fmt"
)

func main() {
	// Create our dealer
	dealer := NewDealer()

	// Create the players
	players := []*Player{NewPlayer(NewHitSoft17()), NewPlayer(NewHitSoft17()),
		NewPlayer(NewSimple()), NewPlayer(NewSimple())}

	var shoe *Shoe
	for i := 0; i < 1_000_000; i++ {
		var newShoe bool
		if shoe == nil || shoe.IsDone() {
			// fmt.Println("New shoe")
			shoe = NewShoe(3)
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
				if action == Hit {
					player.Take(shoe.Pull(), true)
				} else {
					break
				}
			}
		}

		// Dealer goes
		for {
			action := dealer.Action(dealer.TopCard())
			if action == Hit {
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
		wager := player.Wager()
		bankroll := player.Bankroll()
		fmt.Printf("Player(%d - %s) %d : %d = %d%%\n", i, player.Strategy(), wager, bankroll, (bankroll)*100/wager)
	}
}
