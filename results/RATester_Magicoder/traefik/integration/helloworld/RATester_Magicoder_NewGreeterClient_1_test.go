package helloworld

import (
	fmt "fmt"
	"testing"

	grpc "google.golang.org/grpc"
)

func TestNewGreeterClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cc := &grpc.ClientConn{}
	client := NewGreeterClient(cc)

	if client == nil {
		t.Error("Expected client to be not nil")
	}
}
