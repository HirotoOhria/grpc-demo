package main

import (
    "fmt"
    "log"
    "net"
    "os"
    "os/signal"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "image.uploader/api/gen/pb"
    "image.uploader/handler"
)

const port = 50053

func main() {
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
    if err != nil {
        log.Fatalf("failded to listen: %v", err)
    }

    server := grpc.NewServer()

    pb.RegisterImageUploadServiceServer(server, handler.NewImageUploadHandler())
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
