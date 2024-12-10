package model

type PaymentIntentData struct {
	Amount         int64
	Currency       string
	Payment_method string
	Confirm        bool
	Token          string
}