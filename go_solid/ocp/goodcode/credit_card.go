package goodcode

import "fmt"

type CreditCard struct{}

func (c *CreditCard) pay(amount float64) {
	fmt.Println("Processing Payment via CreditCard: ", amount)
}