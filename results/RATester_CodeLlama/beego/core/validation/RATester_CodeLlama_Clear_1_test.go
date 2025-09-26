package validation

import (
	"fmt"
	"testing"
)

func TestClear_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &Validation{}
	v.AddError("key", "message")
	v.Clear()
	if len(v.Errors) != 0 {
		t.Errorf("Expected 0 errors, got %d", len(v.Errors))
	}
}
