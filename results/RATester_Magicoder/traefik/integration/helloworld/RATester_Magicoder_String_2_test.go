package helloworld

import (
	fmt "fmt"
	"testing"
)

func TestString_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	reply := &StreamExampleReply{
		Data: "test data",
	}
	expected := `data:"test data"`
	if got := reply.String(); got != expected {
		t.Errorf("String() = %v, want %v", got, expected)
	}
}
