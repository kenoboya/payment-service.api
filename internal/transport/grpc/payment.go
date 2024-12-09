package grpc_handler

import "payment-api/internal/service"

type PaymentHandler struct {
	services service.Services
}

func NewPaymentHandler(services service.Services) *PaymentHandler{
	return &PaymentHandler{services: services}
}
