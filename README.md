# BlackJack Go
A classic BlackJack card game library in GO


## Purpose
Go package can be used to run a blackjack game. It is also an exercise in Object Oriented Design and 
Test-driven Development

## Classes

### Card
* Name string
* Value int
* *.Display() - Print the name of the Card

### Deck
* Constructor NewDeck() - generate a new deck class
* Cards []Card
* New() - discard current Cards and grab a new deck of card
* Shuffle() - shuffle/randomized the order of Cards
* Draw() - return a single card from the deck and remove it
* Count() - count number of cards remaining in deck

### Hand
* Constructor NewHand()
* Cards []Card
* ReceiveCard(card Card)
* Count()
* Value()
* IsBlackJack()
* Discard()
* Display()

### Player
* Constructor NewPlayer(name string)
* Name string
* Hand Hand
* Hold bool
* ReceiveCard(card *Card)
* Prompt()

### Dealer
* Inherits Player
* Constructor NewDealer()
* DealCard()
* GetNewDeck()

### Game
* Constructor NewGame()
* StartGame()
* NextRound()
* AddPlayer(player Player)
* PlayerCount()
* Evaluate()
* EndGame()
