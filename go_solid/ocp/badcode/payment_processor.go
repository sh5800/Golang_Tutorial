package badcode

import "fmt"

type PaymentProcessor struct{}

func(p *PaymentProcessor) processPayment(paymentMethod string, amount float64) {
	if(paymentMethod == "CreditCard") {
		fmt.Println("Processing Payment via CreditCard: ",amount)
	}else if(paymentMethod == "DebitCard") {
		fmt.Println("Processing Payment via DebitCard: ",amount)
	}else if(paymentMethod == "PayPal"){
		fmt.Println("Processing Payment via PayPal: ",amount)
	}else{
		fmt.Errorf("Unsupported Payment Method")
	}
}