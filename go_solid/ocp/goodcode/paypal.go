package goodcode

import "fmt"

type PayPal struct{}

func (pp *PayPal) Pay(amount float64) {
	fmt.Println("Making Payment via PayPal: ",amount)
}