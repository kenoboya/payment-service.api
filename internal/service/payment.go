package service

import (
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/customer"
	"github.com/stripe/stripe-go/v81/paymentintent"
)

type PaymentsService struct {
}

func NewPaymentsService() *PaymentsService {
	return &PaymentsService{}
}

func (s *PaymentsService) CreatePaymentIntent(customerData model.CustomerData, paymentIntentData model.PaymentIntentData) (transactionID string, err error) {
	paramsCustomer:= &stripe.CustomerParams{
		// Email:
	}
	c, err:= customer.New(paramsCustomer)
	if err != nil {
		// logger
	}

	paramsIntent := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(paymentIntentData.amount),
		Currency: stripe.String(string(stripe.CurrencyUAH)),
		PaymentMethod: stripe.String(paymentIntentData.stripeToken),
		Confirm: stripe.Bool(true),
		Customer: stripe.String(c.ID),
		ReceiptEmail: stripe.String(customerData.Email),
	}
	pi, err:= paymentintent.New(paramsIntent)
	if err != nil{
		// logger
	}

	return pi.ID, nil
}
func (s *PaymentsService) ConfirmPayment(transactionID string) (bool, error) {
	pi, err:= paymentintent.Get(transactionID, nil)
	if err != nil {
		// logger
	}

	if pi.Status == stripe.PaymentIntentStatusSucceeded{
		// logger
		return true, nil
	} else if pi.Status == stripe.PaymentIntentStatusRequiresPaymentMethod || pi.Status == stripe.PaymentIntentStatusRequiresAction {
		// logger
		return false, nil
	} else {
		// logger
		return false, nil
	}
}