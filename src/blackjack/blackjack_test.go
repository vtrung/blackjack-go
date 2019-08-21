package blackjack

import "testing"

//Test Card Class
func TestCard(t *testing.T){
	card := Card{
		"test1",
		1,
	}
	card.Display()
	expname := "test1"
	if card.Name != expname {
		t.Errorf("got %v want %v given", expname, card.Name)
	}
}

//Testing Deck Class
func TestDeck(t *testing.T) {
	// check deck is filled with all cards
	deck := Deck{}
	deck.New()
	if deck.Count() != 52 {
		t.Errorf("got %v want 52 given", deck.Count())
	}

	// check deck draw will remove a card from deck
	card1 := deck.Draw()
	if deck.Count() != 51 {
		t.Errorf("got %v want 51 given", deck.Count())
	}
	// check if correct card was drawn and discarded
	if FindCardInArray(*card1, deck.Cards) {
		t.Errorf("correct card was not discarded from deck")
	}

	// check Shuffle function
	oldorder := deck.Cards
	deck.Shuffle()

	if CardArrayEqual(deck.Cards, oldorder) {
		t.Errorf("Deck failed to shuffle")
		//deck.Display()
	}


}

//Testing Player Class
func TestPlayerDealer(t *testing.T){
	player := NewPlayer("Player1")
	dealer := NewDealer()

	card1 := dealer.DealCard()
	player.ReceiveCard(&card1)

	//Test Player Hand
	if player.Hand.Count() != 1 {
		t.Errorf("Player should have 1 card got %d instead", player.Hand.Count())
	}

	if dealer.Deck.Count() != 51{
		t.Errorf("Dealer Deck should have 51 card got %d instead", dealer.Deck.Count())
	}
	
}

func TestGame(t *testing.T) {
	
	//deck := NewDeck()

	//deck.Display()

}



//Utility Testing Functions
//Check if 2 list of Cards are the same
func CardArrayEqual(a, b []Card) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v.Value != b[i].Value {
			return false
		}
	}
	return true
}

func FindCardInArray(card Card, cards []Card) bool {
	for _, c := range cards{
		if c.Name == card.Name {
			return true
		}
	}
	return false
}