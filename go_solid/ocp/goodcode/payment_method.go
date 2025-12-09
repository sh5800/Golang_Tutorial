package goodcode

type PaymentMethod interface {
	pay(amount float64)
}