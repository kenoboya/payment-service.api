package service

type Services struct {
	Payments Payments
}

type Payments interface {
	CreatePaymentIntent() error
	ConfirmPayment() (bool, error)
}