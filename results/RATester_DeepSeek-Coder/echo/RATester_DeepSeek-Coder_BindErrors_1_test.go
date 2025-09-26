package echo

import (
	"errors"
	"fmt"
	"testing"
)

func TestBindErrors_1(t *testing.T) {
	t.Run("should return nil when no errors have been set", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b := &ValueBinder{}
		got := b.BindErrors()
		if got != nil {
			t.Errorf("expected nil, got %v", got)
		}
	})

	t.Run("should return errors and reset the errors field", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b := &ValueBinder{
			errors: []error{errors.New("error 1"), errors.New("error 2")},
		}
		got := b.BindErrors()
		if len(got) != 2 {
			t.Errorf("expected 2 errors, got %d", len(got))
		}
		if b.errors != nil {
			t.Errorf("expected errors to be reset, got %v", b.errors)
		}
	})
}
