package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"pancake.maker/api/gen/api"
	"pancake.maker/handler"
)

const port = 50051

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed ot listen: %v", err)
	}

	server := grpc.NewServer()
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
