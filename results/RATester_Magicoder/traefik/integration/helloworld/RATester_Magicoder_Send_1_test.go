package helloworld

import (
	fmt "fmt"
	"testing"
)

func TestSend_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	serverStream := &greeterStreamExampleServer{}
	reply := &StreamExampleReply{Data: "test data"}

	err := serverStream.Send(reply)
	if err != nil {
		t.Errorf("Failed to send reply: %v", err)
	}
}
