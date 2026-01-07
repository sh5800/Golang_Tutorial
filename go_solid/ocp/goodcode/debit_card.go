package goodcode

import "fmt"

type DebitCard struct{}

func (c *DebitCard) Pay(amount float64) {
	fmt.Println("Processing Payment via DebitCard: ", amount)
}