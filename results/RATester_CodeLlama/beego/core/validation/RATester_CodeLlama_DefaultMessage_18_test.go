package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_18(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := AlphaDash{}
	want := MessageTmpls["AlphaDash"]
	got := a.DefaultMessage()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
