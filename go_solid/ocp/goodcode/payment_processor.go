package goodcode

type PaymentProcessor struct{}

func(ppr *PaymentProcessor) ProcessPayment(paymentMethod PaymentMethod, amount float64){
	paymentMethod.Pay(amount)
}