package badcode

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

func(i *Invoice) SaveToDatabase(){
	fmt.Println("Invoice saved to database")
}

func(i *Invoice) SendNotification(){
	fmt.Println("Invoice Notification sent successfully")
}