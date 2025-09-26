package helloworld

import (
	fmt "fmt"
	"testing"
)

func TestProtoMessage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	sr := &StreamExampleReply{
		Data: "test",
	}

	sr.ProtoMessage()

	// Add assertion here if needed
}
