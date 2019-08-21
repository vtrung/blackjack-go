package blackjack

import "testing"
import "fmt"

//Test Card Class
func TestCard(t *testing.T){
	fmt.Println("== Testing Card ==")
	card := Card{
		"J of Spades",
		10,
	}
	fmt.Println("** Print Card")
	card.Display()
	expname := "J of Spades"
	if card.Name != expname {
		t.Errorf("got %v want %v given", expname, card.Name)
	}
}

//Testing Deck Class
func TestDeck(t *testing.T) {
	fmt.Println("== Testing Deck ==")
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


func TestHand(t *testing.T){
	fmt.Println("== Testing Hand ==")
	hand := NewHand()
	card1 := Card{
		"J of Spades",
		10,
	}
	hand.ReceiveCard(card1);
	if hand.Count() != 1 {
		t.Errorf("Hand should have 1 card got %d instead", hand.Count())
	}

	hand.Discard()
	if hand.Count() != 0{
		t.Errorf("Hand should have 0 card got %d instead", hand.Count())
	}
}

//Testing Player Class
func TestPlayerDealer(t *testing.T){
	fmt.Println("== Testing Player/Dealer ==")
	player := NewPlayer("Player1")
	dealer := NewDealer()

	card1 := dealer.DealCard()
	player.ReceiveCard(&card1)

	//Test Player Hand
	if player.Hand.Count() != 1 {
		t.Errorf("Player should have 1 card got %d instead", player.Hand.Count())
	}

	if !FindCardInArray(card1, player.Hand.Cards){
		t.Errorf("Player picked up wrong card")
	}

	if dealer.Deck.Count() != 51{
		t.Errorf("Dealer Deck should have 51 card got %d instead", dealer.Deck.Count())
	}

	if FindCardInArray(card1, dealer.Deck.Cards){
		t.Errorf("Dealer Deck discard wrong card")
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