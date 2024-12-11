package service

import (
	"fmt"
	"payment-api/internal/model"
	logger "payment-api/pkg/logger/zap"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/customer"
	"github.com/stripe/stripe-go/v81/paymentintent"
	"go.uber.org/zap"
)

type PaymentsService struct {}

func NewPaymentsService() *PaymentsService {
	return &PaymentsService{}
}

func (s *PaymentsService) CreatePaymentIntent(customerData model.CustomerData, paymentIntentData model.PaymentIntentData) (transactionID string, err error) {
	paramsCustomer := &stripe.CustomerParams{
		Name:  &customerData.Name,
		Phone: &customerData.Phone,
	}

	if customerData.Email != "" {
		paramsCustomer.Email = &customerData.Email
	}

	c, err := customer.New(paramsCustomer)
	if err != nil {
		logger.Error("Failed to create Stripe customer", zap.Error(err))
		return "", fmt.Errorf("failed to create customer: %w", err)
	}

	paramsIntent := &stripe.PaymentIntentParams{
		Amount:       stripe.Int64(paymentIntentData.Amount),
		Currency:     stripe.String(string(stripe.CurrencyUAH)),
		PaymentMethod: stripe.String(paymentIntentData.PaymentMethod),
		Confirm:      stripe.Bool(true),
		Customer:     stripe.String(c.ID),
		ReceiptEmail: stripe.String(customerData.Email),
		Metadata: map[string]string{
			"order_id": fmt.Sprintf("%d", paymentIntentData.OrderID),
		},
	}

	pi, err := paymentintent.New(paramsIntent)
	if err != nil {
		logger.Error("Failed to create PaymentIntent", zap.Error(err))
		return "", fmt.Errorf("failed to create payment intent: %w", err)
	}

	return pi.ID, nil
}

func (s *PaymentsService) ConfirmPayment(transactionID string) (bool, error) {
	pi, err := paymentintent.Get(transactionID, nil)
	if err != nil {
		logger.Error("Failed to get PaymentIntent", zap.Error(err))
		return false, fmt.Errorf("failed to get payment intent: %w", err)
	}

	switch pi.Status {
	case stripe.PaymentIntentStatusSucceeded:
		logger.Info("The transaction has been completed successfully")
		return true, nil
	case stripe.PaymentIntentStatusRequiresPaymentMethod, stripe.PaymentIntentStatusRequiresAction:
		logger.Info("Transaction requires additional action or payment method")
		return false, fmt.Errorf("transaction requires additional action or payment method")
	default:
		logger.Info("Transaction not confirmed")
		return false, fmt.Errorf("transaction not confirmed with status: %s", pi.Status)
	}
}
