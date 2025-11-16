package main

type Card struct{
	cardNumber string
	userName string
}

func(c *Card) NewCard(CardNumber, UserName string){
	c.cardNumber = CardNumber
	c.userName = UserName
}

func(c *Card) GetCardNumber() string{
	return c.cardNumber
}

func (c *Card) GetUserName() string{
	return c.userName
}