package handler

import (
	"context"
	"fmt"
	"github.com/greeflas/itea_lessons/pb"
)

type GreetingServiceHandler struct {
	pb.UnimplementedGreetingServiceServer
}

func NewGreetingServiceHandler() *GreetingServiceHandler {
	return &GreetingServiceHandler{}
}

func (h *GreetingServiceHandler) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	name := "Guest"

	if req.Name != "" {
		name = req.Name
	}

	message := fmt.Sprintf("Hello, %s!", name)

	return &pb.SayHelloResponse{
		Message: message,
	}, nil
}
