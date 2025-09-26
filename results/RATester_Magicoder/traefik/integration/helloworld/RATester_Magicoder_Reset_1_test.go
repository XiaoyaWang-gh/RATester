package helloworld

import (
	fmt "fmt"
	"testing"
)

func TestReset_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	reply := &StreamExampleReply{
		Data: "test data",
	}

	reply.Reset()

	if reply.Data != "" {
		t.Errorf("Expected Data to be empty after Reset, but got: %s", reply.Data)
	}
}
