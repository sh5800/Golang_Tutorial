package main

type PaymentService struct{
	paymentMethods map[string]PaymentMethod
}

func (ps *PaymentService) NewPaymentService (PaymentMethods map[string]PaymentMethod){
	ps.paymentMethods = PaymentMethods
}

func (ps *PaymentService) AddPayment(Name string, PayMethod PaymentMethod){
	ps.paymentMethods[Name] = PayMethod
}

func (ps *PaymentService) MakePayment(Name string){
	pm := ps.paymentMethods[Name]
	pm.Pay() //run time polymorphism
}