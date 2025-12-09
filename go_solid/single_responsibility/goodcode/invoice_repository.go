package goodcode

import "fmt"

type InvoiceRepository struct{}
func(ir *InvoiceRepository) SaveToDatabase(i *Invoice){
	fmt.Println("Invoice saved to database")
}