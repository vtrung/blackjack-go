package blackjack

import (  
    "fmt"
)

type Card struct{
	Name string
	Value int
}

func (c Card) Display(){
	fmt.Println(c.Name)
}