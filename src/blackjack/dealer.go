package blackjack

type Dealer struct{
	Player *Player
	Deck Deck
}

func NewDealer() *Dealer {
	deck := NewDeck()
	deck.Shuffle()
	player := NewPlayer("Dealer")
	dealer := Dealer{
		Player: player,
		Deck: *deck,
	}

	return &dealer
}

func (d *Dealer) DealCard() Card{
	card := d.Deck.Draw()
	return *card
}

func (d *Dealer) GetNewDeck(){
	deck := NewDeck()
	deck.Shuffle()
	d.Deck = *deck
}

