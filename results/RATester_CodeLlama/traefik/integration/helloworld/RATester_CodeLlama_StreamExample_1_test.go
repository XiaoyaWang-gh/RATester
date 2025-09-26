package helloworld

import (
	context "context"
	fmt "fmt"
	"io"
	"testing"

	grpc "google.golang.org/grpc"
)

func TestStreamExample_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	c := &greeterClient{}
	in := &StreamExampleRequest{}
	opts := []grpc.CallOption{}
	stream, err := c.StreamExample(ctx, in, opts...)
	if err != nil {
		t.Fatal(err)
	}
	if err := stream.CloseSend(); err != nil {
		t.Fatal(err)
	}
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
	}
}
