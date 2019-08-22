# BlackJack Go
A classic BlackJack card game library in GO


## Purpose
Go package can be used to run a blackjack game. It is also an exercise in Object Oriented Design and 
Test-driven Development

## Usage
add: import "github.com/vtrung/blackjack-go/src/blackjack"

## Testing
run in package directory: 'go test' <br/>
see blackjack_test.go for testing parameters

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
* ReceiveCard(card Card) - Recieve card and place it in hand
* Count() - return int. Count of card in hand
* Value() - return int. Highest Blackjack hand value
* IsBlackJack() - return bool.  Is this hand a blackjack
* Discard() - drop cards
* Display() - prints cards in hand

### Player
* Constructor NewPlayer(name string)
* Name string
* Hand Hand
* Hold bool
* ReceiveCard(card *Card) - take card and place in hand
* Prompt() - print users cards

### Dealer
* Inherits Player
* Constructor NewDealer()
* DealCard() - return Card. draw a card from the deck and return it
* GetNewDeck() - discard current deck and get a new shuffled deck

### Game
* Constructor NewGame()
* StartGame() - Start first phase of blackjack game, pass 2 cards to each player and dealer
* NextRound() - Print cards of each player
* AddPlayer(player Player) - Add player to game
* PlayerCount() - return int. number of players in game
* Evaluate() - evaluate value of all players and dealer and determine win/lose/draw
* EndGame() - End game by discard hands and grabbing new deck
