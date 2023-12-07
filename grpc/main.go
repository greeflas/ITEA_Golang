package main

import (
	"context"
	"github.com/greeflas/itea_lessons/handler"
	"github.com/greeflas/itea_lessons/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}

	greetingServiceHandler := handler.NewGreetingServiceHandler()

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(ServerLoggingInterceptor),
	)

	pb.RegisterGreetingServiceServer(grpcServer, greetingServiceHandler)

	reflection.Register(grpcServer)

	log.Println("Starting gRPC server...")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func ServerLoggingInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	log.Printf("incoming request: %s", info.FullMethod)

	return handler(ctx, req)
}
