package validation

import (
	"fmt"
	"testing"
)

func TestHasErrors_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &Validation{}

	// Test case 1: No errors
	v.Clear()
	if v.HasErrors() {
		t.Error("Expected HasErrors to return false, but it returned true")
	}

	// Test case 2: Errors present
	v.AddError("field", "error message")
	if !v.HasErrors() {
		t.Error("Expected HasErrors to return true, but it returned false")
	}
}
