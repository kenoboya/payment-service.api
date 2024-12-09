package model

type PaymentIntentData struct {
	StripeToken string
	Amount      int64
}