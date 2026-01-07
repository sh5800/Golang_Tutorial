package goodcode

type PaymentMethod interface {
	Pay(amount float64)
}