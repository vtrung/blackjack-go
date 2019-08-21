package blackjack

type Hand struct{
	Cards []Card
}

func NewHand() *Hand{
	cards := make([]Card, 0)
	hand := Hand{ Cards: cards}
	return &hand
}

func (h *Hand) ReceiveCard(card Card){
	h.Cards = append(h.Cards, card)
}

func (h *Hand) Value() int {
	sum := 0
	aces := 0
	for _,c := range h.Cards{
		sum += c.Value
		if(c.Value == 11){
			aces++
		}
	}
	// Convert ace value to 1 if sum is greater than 21
	for sum > 21 && aces > 0{
		sum -= 10
		aces--
	}

	return sum
}

func (h *Hand) Count() int {
	return len(h.Cards)
}

func (h *Hand) IsBlackJack() bool {
	if h.Count() == 2 && h.Value() == 21 {
		return true
	}
	return false
}

func (h *Hand) IsOver() bool{
	return (h.Value() > 21)
}

func (h *Hand) Display(){
	for _, c := range h.Cards{
		c.Display()
	}
}

func (h *Hand) Discard(){
	h.Cards = make([]Card, 0)
}