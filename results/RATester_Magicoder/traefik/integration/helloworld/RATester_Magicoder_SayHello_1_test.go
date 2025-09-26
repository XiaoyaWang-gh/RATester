package helloworld

import (
	context "context"
	fmt "fmt"
	"testing"

	grpc "google.golang.org/grpc"
)

func TestSayHello_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	client := &greeterClient{
		cc: &grpc.ClientConn{},
	}

	req := &HelloRequest{
		Name: "John Doe",
	}

	reply, err := client.SayHello(ctx, req)
	if err != nil {
		t.Fatalf("SayHello() failed: %v", err)
	}

	if reply.GetMessage() != "Hello John Doe" {
		t.Errorf("SayHello() = %q, want %q", reply.GetMessage(), "Hello John Doe")
	}
}
