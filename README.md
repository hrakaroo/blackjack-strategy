# BlackJack Strategy

## Overview

I'm not a big gambler in general, but when I do go to Las Vegas I usually play 1/2 no-limit Poker.  I find the
dynamics of the game play to be interesting and with a fairly conservative strategy I'm usually in the positive
after most trips.  (But not enough that I'm ready to quit my day job.)

However, despite sitting at a table with eight other people, it's not a hugely social game.  If I'm with friends 
I'll usually play BlackJack with them, which I'm not super great at.  And while there exist many BlackJack
_trainer_ apps, they mostly require rote memorization of charts which is not terribly fun.  So instead I 
thought it might be interesting to wire up some simulations and see what actually impacts your odds.

This is a side project, so I'm leaning heavily to clarity over speed.  I don't really care how long the simulations
take to execute, so long as I can come back to my code in a week/month/year and easily understand what I wrote.

## Approach

First, I just need to simulate a card, deck, shoe and
the basic play action.  After that is complete want to 
figure out some percentages over 10,000+ simulations and
then start to play with different strategies.

## Results so far

The results are calculated by keeping a running total of how much money we have bet
vs what we have won.  A result of 100% would mean that we basically broke even.
Values less than 100% indicates the strategy loses money and values over 100% indicates
the strategy wins money.  

```
Player(0) - Hit Soft 17          - 2000000 : 1879993 = 94.00%
Player(1) - No Split/No Double   - 2000000 : 1937684 = 96.88%
Player(2) - No Split             - 2204542 : 2174714 = 98.65%
```

### Hit Soft 17

This is basically the dealers strategy where you hit on anything below a 17 or if the
17 is a soft 17 (ie, one of the cards is an Ace). (When I was younger and didn't understand
all the rules of blackjack I couldn't understand how, if you were matching the dealers actions,
this wasn't a break even strategy.)

#### No Split/No Double

This is the "perfect" strategy from 
   https://www.blackjackapprenticeship.com/blackjack-strategy-charts/

but without the ability to double down or split. 

#### No Split

Basically the same as the previous one, but allows doubling down on bets.



