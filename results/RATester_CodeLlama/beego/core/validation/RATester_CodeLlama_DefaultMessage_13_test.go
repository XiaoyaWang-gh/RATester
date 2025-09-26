package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := Email{}
	want := MessageTmpls["Email"]
	got := e.DefaultMessage()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
