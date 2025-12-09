package goodcode

import "fmt"
type Invoice struct {
	Amount float64
}

func (i *Invoice) NewInvoice(Amount float64) {
	i.Amount = Amount
}

func(i *Invoice) GenerateInvoice(){
	fmt.Println("Invoice Generated for amount:",i.Amount)
}

