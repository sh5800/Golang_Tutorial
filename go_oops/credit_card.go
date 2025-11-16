package main

import "fmt"
type CreditCard struct{
	Card
}

func (cc *CreditCard) NewCreditCard(CardNumber, UserName string){
	cc.Card.NewCard(CardNumber, UserName)
}

func (cc *CreditCard) Pay(){
	fmt.Println("Paying with credit card: ",cc.Card.GetCardNumber())
}