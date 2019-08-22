package blackjack

type Player struct{
	Name string
	Hand Hand
	Hold bool
}

func NewPlayer(name string) *Player{
	hand := *NewHand()
	player := Player{Name: name, Hand: hand, Hold: false}
	return &player
}

func(p *Player) ReceiveCard(card Card){
	p.Hand.ReceiveCard(card)
}

func(p *Player) Prompt(){
	p.Hand.Display()
}

func(p *Player) Discard(){
	p.Hand = *NewHand()
}