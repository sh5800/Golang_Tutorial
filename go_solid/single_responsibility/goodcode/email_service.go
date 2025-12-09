package goodcode

import "fmt"

type InvoiceNotifier struct{
	
}
func(in *InvoiceNotifier) SendNotification(i *Invoice){
	fmt.Println("Invoice Notification sent successfully")
}