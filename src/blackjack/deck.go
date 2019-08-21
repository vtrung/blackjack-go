
package blackjack

import "math/rand"
import "fmt"
import "time"

type Deck struct {
	Cards []Card
}

func NewDeck() *Deck{
	deck:= Deck{}
	deck.New()
	return &deck
}

//Draw a new deck of cards
func (d *Deck) New() {
	cards := make([]Card, 0)
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	suits := []string{"Diamonds", "Clubs", "Hearts", "Spades"}
	for i, v := range values {
		for _, suit := range suits {
			if i > 8 {
				cvalue := 10
				if i > 11 {
					cvalue = 11
				}
				card := Card{v + " of " + suit, cvalue}
				cards = append(cards, card)
			} else {
				card := Card{v + " of " + suit, i + 2}
				cards = append(cards, card)
			}

		}
	}
	d.Cards = cards
}

// randomized the order the the current deck
func (d *Deck) Shuffle() {
	length := len(d.Cards)
	if length > 0 {
		rand.Seed(time.Now().UTC().UnixNano())
		r := rand.New(rand.NewSource(time.Now().Unix()))
		
		newslice := make([]Card, length)

		m := make([]int, length)
		for i := 0; i < length; i++ {
			j := r.Intn(i + 1)
			m[i] = m[j]
			m[j] = i
		}

		for i, randIndex := range m {
			newslice[i] = d.Cards[randIndex]
		}

		fmt.Println(newslice)
		d.Cards = newslice
	}
}

// count number of cards remaining
func (d *Deck) Count() int{
	return len(d.Cards)
}

// draw a card from the deck and remove it
func (d *Deck) Draw() *Card {
	//draw top
	//card1 := Card{"1 Diamond", 1}
	//POP card from end
	card1 := d.Cards[len(d.Cards)-1]
	d.Cards = d.Cards[:len(d.Cards)-1]
	return &card1
	//return nil
}

func (d *Deck) Display() {
	for _, card := range d.Cards {
		card.Display()
	}
}
