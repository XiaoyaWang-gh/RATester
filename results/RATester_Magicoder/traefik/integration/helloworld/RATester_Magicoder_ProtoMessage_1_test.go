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

	reply := &StreamExampleReply{
		Data: "test data",
	}

	// Call the method under test
	reply.ProtoMessage()

	// Assert that the method under test does not panic or return an error
	// If the method under test does not return an error, the test will pass
}
