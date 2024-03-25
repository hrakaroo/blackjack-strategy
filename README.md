# BlackJack Strategy

## Overview

I'm not a big gambler in general, but when I do go to Las Vegas I usually play 1/2 no-limit Poker.  I find the
dynamics of the game play to be interesting and with a fairly conservative strategy I'm usually in the positive
after most trips.  (But not enough that I'm ready to quit my regular job and play full time.)

However, despite sitting at a table with 8-9 other people, it's not a hugely social game.  If I'm in 
Las Vegas with friends (unless they are my hard core Poker friends) I will usually play BlackJack with them
instead.  And while there exist many BlackJack _trainer_ apps, they mostly require rote memorization of charts
which is not terribly fun.  So instead I thought it might be interesting to wire up some simulations and see if I could emperically generate the solver charts and maybe, along the way, start to internalize the logic.

This is a side project, so I'm heavily leaning to clarity over speed.  I don't really care how long they take to execute, so long as I can come back to my code in a week/month/year and easily understand what I wrote.

## Phases

First, I just need to simulate a card, deck, shoe and
the basic play action.  After that is complete want to 
figure out some percentages over 10,000+ simulations and
then start to play with different strategies.

Somewhere along the way I'll probably want to see about
incorporating a drawing library so I can generate some of
the solver charts and see how they match up to the 
others I've seen posted.

## Results

The results are calculated by keeping a running total of how much money we have bet
vs what we have won.  So we have a running `wagers` total that gets incremented _on every bet_.
Then we have a `wins` which gets incremented at the end of every hand.  

```
Player(0 - Hit Soft 17) 2000000 : 1878352 = 93.92%
Player(1 - Simple 1) 2000000 : 1931515 = 96.58%
Player(2 - Simple 2) 2206512 : 2175779 = 98.61%
```

### Hit Soft 17

This is basically the dealers strategy where you hit on anything below a 17 or if the
17 is a soft 17 (ie, one of the cards is an Ace)

#### Simple 1

This is the "perfect" strategy from 
   https://www.blackjackapprenticeship.com/blackjack-strategy-charts/

but without the ability to double down or split. 

#### Simple 2

Basically the same as Simple1, but allows double down.



