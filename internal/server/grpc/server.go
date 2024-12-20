package grpc_server

import (
	"net"
	"payment-api/internal/config"
	grpc_handler "payment-api/internal/transport/grpc"
	logger "payment-api/pkg/logger/zap"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type server struct {
	srv *grpc.Server
	addr string
	handler *grpc_handler.PaymentHandler
}

func NewServer(config config.GrpcConfig, handler *grpc_handler.PaymentHandler) (*server){
	return &server{
		srv: grpc.NewServer(),
		addr: config.Addr,
		handler: handler,
	}
}

func (s *server) Run() error{
	lis, err:= net.Listen("tcp", s.addr)
	if err != nil{
		logger.Error("Failed to listen on TCP",
			zap.String("server", "grpc"),
			zap.Error(err),
		)
		return err
	}
	logger.Info("Starting gRPC server",
		zap.String("address", s.addr),
	)
	
	// ... TODO proto.Payment
	
	return nil
}

func (s *server) Stop() {
	s.srv.GracefulStop()
	logger.Info("gRPC server has stopped gracefully",
		zap.String("server", "grpc"),
	)
}