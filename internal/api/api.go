package api

import (
	"os"
	"os/signal"
	"payment-api/internal/config"
	grpc_server "payment-api/internal/server/grpc"
	"payment-api/internal/service"
	grpc_handler "payment-api/internal/transport/grpc"
	logger "payment-api/pkg/logger/zap"
	"syscall"

	"go.uber.org/zap"
)

func Run(configDIR string, envDIR string) {
	logger.InitLogger()
	cfg, err := config.Init(configDIR, envDIR)
	if err != nil {
		logger.Fatal("Failed to initialize config",
			zap.Error(err),
			zap.String("context", "Initializing application"),
			zap.String("version", "1.0.0"),
			zap.String("environment", "development"),
		)
	}

	services := service.NewServices(service.NewPaymentsService())
	handler := grpc_handler.NewPaymentHandler(services)

	// GRPC
	grpcServer := grpc_server.NewServer(cfg.GRPC, handler)
	go func() {
		if err := grpcServer.Run(); err != nil {
			logger.Fatalf("The grpc server didn't start: %s\n", err)
		}
	}()

	// EXIT
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	grpcServer.Stop()
}