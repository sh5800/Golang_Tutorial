package main

import "fmt"
type DebitCard struct{
	Card
}

func(dc *DebitCard) NewDebitCard(CardNumber, UserName string){
	dc.Card.NewCard(CardNumber,UserName)
}

func(dc *DebitCard) Pay(){
	fmt.Println("Paying with debit card: ",dc.Card.GetCardNumber())
}