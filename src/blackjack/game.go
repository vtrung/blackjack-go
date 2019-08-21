package blackjack

import "fmt"

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

func (g *Game) StartGame(){
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
	card2.Display()
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
	dealerBlackJack := g.Dealer.Player.Hand.IsBlackJack()
	for !dealerBlackJack && g.Dealer.Player.Hand.Value() < 17 {
		card := g.Dealer.DealCard()
		g.Dealer.Player.ReceiveCard(&card)
		card.Display()
	}
	

	for _, p := range g.Players{
		fmt.Println(p.Name)
		p.Hand.Display()
		if p.Hand.IsBlackJack(){
			if dealerBlackJack{
				fmt.Println("Draw")
			} else {
				fmt.Println("Win")
			}
		} else if dealerBlackJack || g.Dealer.Player.Hand.Value() > p.Hand.Value(){
			fmt.Println("Lose")
		} else {
			fmt.Println("Win")
		}
	}
}

func (g *Game) EndGame(){
	g.Dealer.Player.Hand.Discard()
	for _, p := range g.Players{
		p.Hand.Discard()
	}
}