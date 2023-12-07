package handler

import (
	"context"
	"github.com/greeflas/itea_lessons/pb"
	"testing"
)

func TestSayHello(t *testing.T) {
	req := &pb.SayHelloRequest{
		Name: "Tester",
	}

	handler := NewGreetingServiceHandler()
	res, err := handler.SayHello(context.Background(), req)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedMessage := "Hello, Tester!"

	if res.Message != expectedMessage {
		t.Errorf("invalid message: got: %s, want: %s", res.Message, expectedMessage)
	}
}
