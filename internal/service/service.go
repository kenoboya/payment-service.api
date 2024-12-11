package service

import "payment-api/internal/model"

type Services struct {
	Payments Payments
}

func NewServices(paymentsService Payments) *Services {
	return &Services{Payments: paymentsService}
}
type Payments interface {
	CreatePaymentIntent(customerData model.CustomerData, paymentIntentData model.PaymentIntentData) (transactionID string, err error)
	ConfirmPayment(transactionID string) (bool, error)
}