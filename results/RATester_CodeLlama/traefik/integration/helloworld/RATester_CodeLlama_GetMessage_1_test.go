package helloworld

import (
	fmt "fmt"
	"testing"
)

func TestGetMessage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &HelloReply{}
	if got := m.GetMessage(); got != "" {
		t.Errorf("GetMessage() = %v, want %v", got, "")
	}
}
