package grpc_handler

import (
	"context"
	"payment-api/internal/model"
	"payment-api/internal/server/grpc/proto"
	"payment-api/internal/service"
	logger "payment-api/pkg/logger/zap"

	"go.uber.org/zap"
)

type PaymentHandler struct {
	services *service.Services
	proto.UnimplementedPaymentServiceServer
}

func NewPaymentHandler(services *service.Services) *PaymentHandler{
	return &PaymentHandler{services: services}
}

func (h *PaymentHandler) CreatePaymentIntent(ctx context.Context, req *proto.PaymentRequest) (*proto.PaymentResponse, error) {
	transactionID, err:= h.services.Payments.CreatePaymentIntent(model.CustomerData{
		Name: req.Customer.Name,
		Email: req.Customer.Email,
		Phone: req.Customer.Phone,
	}, model.PaymentIntentData{
		Amount: req.PaymentIntent.Amount ,
		Currency: req.PaymentIntent.Currency,
		PaymentMethod: req.PaymentIntent.PaymentMethod, 
	    Confirm: req.PaymentIntent.Confirm   ,    
	    OrderID:       req.PaymentIntent.OrderId,
	})
	if err != nil {
		logger.Error(
			zap.String("action", "CreatePaymentIntent()"),
			zap.Error(err),
		)
		return &proto.PaymentResponse{}, err
	}
	return &proto.PaymentResponse{TransactionId: transactionID}, nil
}