package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := Phone{}
	want := MessageTmpls["Phone"]
	got := p.DefaultMessage()
	if got != want {
		t.Errorf("Phone.DefaultMessage() = %v, want %v", got, want)
	}
}
