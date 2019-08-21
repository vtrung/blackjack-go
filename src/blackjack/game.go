package blackjack

type Game struct{
	Dealer Dealer
	Players []Player
}

func NewGame() *Game{
	dealer := NewDealer()
	players := make([]Player, 0)
	game := Game{
		Dealer: *dealer, 
		Players: players}
	return &game
}

func (g *Game) Start(){
	for _, p := range g.Players {
		card1 := g.Dealer.DealCard()
		card2 := g.Dealer.DealCard()
		p.ReceiveCard(&card1)
		p.ReceiveCard(&card2)
	}
	card1 := g.Dealer.DealCard()
	card2 := g.Dealer.DealCard()
	g.Dealer.Player.ReceiveCard(&card1)
	g.Dealer.Player.ReceiveCard(&card2)
}

func (g *Game) NextRound(){
	for _, p := range g.Players {
		p.Prompt()
	}
}



func (g *Game) AddPlayer(player Player){
	g.Players = append(g.Players, player)
}

func (g *Game) PlayerCount() int{
	return len(g.Players)
}

func (g *Game) Evaluate() {
	g.Dealer.Player.Hand.Display()
	for g.Dealer.Player.Hand.Value() < 17 {
		card := g.Dealer.DealCard()
		g.Dealer.Player.ReceiveCard(&card)
	}
}