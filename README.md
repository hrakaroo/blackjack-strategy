# BlackJack Strategy

## Overview

I'm not a big gambler in general, and when I do go to Las Vegas I usually play 1/2 no-limit Poker.  I find the
dynamics of the game play interesting, and with a fairly conservative strategy I'm usually in the positive
after most trips.  (But not enough that I'm ready to quit my day job.)

However, despite sitting at a table with eight other people, it's not a hugely social game.  If I'm with friends 
I'll usually play BlackJack with them which, I'll be honest, I'm not terribly good at.  And while there exist 
many BlackJack _trainer_ apps, most of them require a fair bit of rote memorization of charts which is not 
terribly fun.  So instead I thought it might be interesting to wire up some simulations and see what actually 
impacts your odds and maybe along the way I'll start to internalize some of these rules.

This is a side project, so I'm leaning heavily to clarity over speed.  I don't really care how long the simulations
take to execute, so long as I can come back to my code in a week/month/year and easily understand what I wrote.

## Approach

First, I just need to simulate a card, deck, shoe and
the basic play action.  After that is complete want to 
figure out some percentages over 1,000,000 simulations and
then start to play with different strategies.

## Results so far

The results are calculated by keeping a running total of how much money we have bet
vs what we have won.  A result of 100% would mean that we basically broke even.
Values less than 100% indicates the strategy loses money and values over 100% indicates
the strategy wins money.  It should come as no suprise that without card counting 
we are under 100%.

```
Player(0) - Hit Soft 17          - 2000000 : 1878457 = 93.92%
Player(1) - Never Bust           - 2168178 : 2069800 = 95.46%
Player(2) - No Split/No Double   - 2000000 : 1934630 = 96.73%
Player(3) - No Split             - 2205580 : 2176473 = 98.68%
Player(4) - Perfect              - 2269934 : 2263038 = 99.70%
```

### Hit Soft 17

This is basically the dealers strategy where the player hits on anything below a 17, or on 17 if 
it's a soft 17 (ie, one of the cards is an Ace).

### Never Bust

Simple strategy where the player is betting the dealer will bust so never hits on anything
above a hard 12. 

This initially had a 93% success rate until it was updated to allow doubling on a
10 or 11.  That one change increased the success rate to 95.5%.

### No Split/No Double

This is the "perfect" strategy from 
   https://www.blackjackapprenticeship.com/blackjack-strategy-charts/

but without the ability to double down or split. Mostly because this was the easiest
to simulate first, but also because I'm curious how much each of these strategies
improves your odds.

### No Split

Basically the same as the previous one, but allows doubling down on bets.  Again, 
adding doubling adds on two percentage points.

#### Perfect

This is basically the full implementation of the perfect strategy as defined above.
This takes into account all splits and doubles.  With a perfect strategy, over a long enough
period of time, you can achive an average 99.7%.  So still losing, but only by a small
fraction, but you need to stick to the plan.

### Card Counting

Okay, now we get into some fun stuff.  Full card counting is not trivial.  It requires a fair 
bit of concentration, but given we are at 99.7% with the perfect strategy, I'm curious how much
counting you would need to do to push that number over the edge.  For example, what if you just
counted the last n hands?  How many hands would you need to count until you were even or slightly
positive EV.  And when you do full card counting, what is your percentage like then?

For card counting we are going to use the basic strategy:

2-6 = +1<br>
7-9 = 0<br>
10-A = -1<br>

When we are in the positive by more than `k` we can increase our bets.