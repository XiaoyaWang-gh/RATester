package validation

import (
	"fmt"
	"testing"
)

func TestString_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Error{
		Message: "test",
	}
	if e.String() != "test" {
		t.Errorf("String() = %v, want %v", e.String(), "test")
	}
}
