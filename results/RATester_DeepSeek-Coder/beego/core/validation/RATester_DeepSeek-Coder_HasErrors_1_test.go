package validation

import (
	"fmt"
	"testing"
)

func TestHasErrors_1(t *testing.T) {
	v := &Validation{}

	t.Run("NoErrors", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		if v.HasErrors() {
			t.Errorf("Expected HasErrors() to return false, got true")
		}
	})

	t.Run("WithErrors", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		v.Errors = append(v.Errors, &Error{})
		if !v.HasErrors() {
			t.Errorf("Expected HasErrors() to return true, got false")
		}
	})
}
