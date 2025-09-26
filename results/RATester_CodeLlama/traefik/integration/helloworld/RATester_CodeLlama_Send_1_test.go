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

	x := &greeterStreamExampleServer{}
	m := &StreamExampleReply{}
	if err := x.Send(m); err != nil {
		t.Fatalf("Send() = %v, want nil", err)
	}
}
