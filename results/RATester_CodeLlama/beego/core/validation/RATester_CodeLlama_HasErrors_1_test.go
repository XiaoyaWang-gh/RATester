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
	v.Errors = []*Error{
		{
			Message: "error1",
		},
		{
			Message: "error2",
		},
	}
	if !v.HasErrors() {
		t.Errorf("HasErrors() = %v, want %v", v.HasErrors(), true)
	}
}
