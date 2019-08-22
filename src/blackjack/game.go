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
	for i, _ := range g.Players {
		card1 := g.Dealer.DealCard()
		card2 := g.Dealer.DealCard()
		//p.ReceiveCard(card11)
		g.Players[i].ReceiveCard(card1)
		g.Players[i].ReceiveCard(card2)
		//p.ReceiveCard(card21)
	}
	card1 := g.Dealer.DealCard()
	card2 := g.Dealer.DealCard()
	g.Dealer.Player.ReceiveCard(card1)
	g.Dealer.Player.ReceiveCard(card2)
	fmt.Printf("\nGame Start\n\n")
	fmt.Println("Dealer last card")
	card2.Display()
}

func (g *Game) NextRound(){
	fmt.Printf("\nStatus\n\n")
	for _, p := range g.Players {
		println(p.Name)
		p.Prompt()
	}
}

func (g *Game) AddPlayer(player *Player){
	g.Players = append(g.Players, *player)
}

func (g *Game) PlayerCount() int{
	return len(g.Players)
}

func (g *Game) Evaluate() {
	fmt.Printf("\nEvaluation Game\n\n")
	fmt.Println("==Dealer===")
	for _, c := range g.Dealer.Player.Hand.Cards{
		fmt.Println(c.Name)
	}

	dealerBlackJack := g.Dealer.Player.Hand.IsBlackJack()
	for !dealerBlackJack && g.Dealer.Player.Hand.Value() < 17 {
		card := g.Dealer.DealCard()
		g.Dealer.Player.ReceiveCard(card)
		fmt.Println("Draw - " + card.Name)
	}
	fmt.Printf("Dealer Value: %v\n", g.Dealer.Player.Hand.Value())
	dealerOver := g.Dealer.Player.Hand.IsOver()
	if dealerOver {
		fmt.Println("Dealer Over")
	}
	
	fmt.Println("== Players ===")

	for _, p := range g.Players{
		fmt.Println(p.Name)
		for _, c := range p.Hand.Cards{
			fmt.Println(c.Name)
		}
		fmt.Printf("Value: %v\n", p.Hand.Value())
		if p.Hand.IsOver(){
			fmt.Println("Lose")
		} else if dealerOver{
			fmt.Println("Win")
		} else if p.Hand.IsBlackJack(){
			if dealerBlackJack{
				fmt.Println("Draw")
			} else {
				fmt.Println("Win")
			}
		} else if dealerBlackJack || g.Dealer.Player.Hand.Value() > p.Hand.Value(){
			fmt.Println("Lose")
		} else if g.Dealer.Player.Hand.Value() == p.Hand.Value(){
			fmt.Println("Draw")
		} else {
			fmt.Println("Win")
		}
		fmt.Println()
	}
}

func (g *Game) EndGame(){
	g.Dealer.Player.Discard()
	for i, _ := range g.Players{
		g.Players[i].Discard()
	}
	g.Dealer.GetNewDeck()
}