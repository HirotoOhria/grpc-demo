package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"pancake.maker/api/gen/api"
	"pancake.maker/handler"
)

const port = 50051

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	grpc_zap.ReplaceGrpcLoggerV2(logger)

	server := grpc.NewServer(genInterceptor(logger))
	api.RegisterPancakeBakerServiceServer(server, handler.NewBakerHandler())
	// To debug on grpc_cli.
	reflection.Register(server)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		server.Serve(lis)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	server.GracefulStop()
}

func genInterceptor(logger *zap.Logger) grpc.ServerOption {
	return grpc.UnaryInterceptor(
		grpc_zap.UnaryServerInterceptor(logger),
	)
}
