package model

type PaymentIntentData struct {
	Amount        int64
	Currency      string
	PaymentMethod string
	Confirm       bool
	OrderID       int64
}