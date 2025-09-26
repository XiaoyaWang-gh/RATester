package gin

import (
	"fmt"
	"testing"
)

func TestSetType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	msg := &Error{}
	msg.SetType(ErrorType(1))
	if msg.Type != ErrorType(1) {
		t.Errorf("msg.Type = %v, want %v", msg.Type, ErrorType(1))
	}
}
