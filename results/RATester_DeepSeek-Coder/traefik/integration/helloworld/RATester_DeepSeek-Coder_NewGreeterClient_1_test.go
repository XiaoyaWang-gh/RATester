package helloworld

import (
	context "context"
	fmt "fmt"
	"testing"
	"time"

	grpc "google.golang.org/grpc"
)

func TestNewGreeterClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn, err := grpc.DialContext(context.Background(), "test", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := NewGreeterClient(conn)
	if client == nil {
		t.Fatalf("NewGreeterClient returned nil")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = client.SayHello(ctx, &HelloRequest{Name: "Test"})
	if err != nil {
		t.Errorf("SayHello failed: %v", err)
	}
}
