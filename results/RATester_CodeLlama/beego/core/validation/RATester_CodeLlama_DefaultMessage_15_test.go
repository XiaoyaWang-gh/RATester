package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := Enum{
		Rules: "Enum",
		Key:   "Enum",
	}
	want := "Enum Requires that the field must be within the enumerated value"
	got := e.DefaultMessage()
	if got != want {
		t.Errorf("DefaultMessage() = %q, want %q", got, want)
	}
}
