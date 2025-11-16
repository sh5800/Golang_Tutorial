package main

func main(){
	ps := new(PaymentService)
	ps.NewPaymentService(make(map[string]PaymentMethod))

	upi := new(UPI)
	upi.NewUPI("shreyashkashyap05")
	ps.AddPayment("upi",upi)

	debitCard := new(DebitCard)
	debitCard.NewDebitCard("1234","SHREYASH KASHYAP")
	ps.AddPayment("debit_card",debitCard)

	creditCard := new(CreditCard)
	creditCard.NewCreditCard("5678","SHREYASH KASHYAP")
	ps.AddPayment("credit_card",creditCard)

	wallet := new(Wallet)
	ps.AddPayment("wallet",wallet)

	ps.MakePayment("upi")
	ps.MakePayment("debit_card")
	ps.MakePayment("credit_card")
	ps.MakePayment("wallet")
}