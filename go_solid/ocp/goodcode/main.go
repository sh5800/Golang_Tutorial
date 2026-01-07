package goodcode

func main(){
	processor := new(PaymentProcessor)
	paymentMethod := new(CreditCard)

	processor.ProcessPayment(paymentMethod,100.00)
}