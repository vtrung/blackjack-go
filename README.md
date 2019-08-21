# BlackJack
Creating the classic BlackJack card game library in GO

## Classes

### Card
Name string
Value int
*.Display() - Print the name of the Card

### Deck
Cards []Card
NewDeck() - generate a new deck class
*.New() - discard current Cards and grab a new deck of card
*.Shuffle() - shuffle/randomized the order of Cards
*.Draw() - return a single card from the deck and remove it
*.Count() - count number of cards remaining in deck

### Hand
Cards []Card

### Player

### Dealer

### Game
