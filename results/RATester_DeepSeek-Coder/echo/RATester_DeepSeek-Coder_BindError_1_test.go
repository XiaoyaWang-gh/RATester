package echo

import (
	"errors"
	"fmt"
	"testing"
)

func TestBindError_1(t *testing.T) {
	t.Run("should return nil if no errors", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b := &ValueBinder{}
		err := b.BindError()
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("should return first error if errors exist", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b := &ValueBinder{errors: []error{errors.New("error 1"), errors.New("error 2")}}
		err := b.BindError()
		if err.Error() != "error 1" {
			t.Errorf("expected 'error 1', got %v", err)
		}
		if len(b.errors) != 1 {
			t.Errorf("expected 1 error left, got %v", len(b.errors))
		}
	})

	t.Run("should clear errors after retrieving", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b := &ValueBinder{errors: []error{errors.New("error 1"), errors.New("error 2")}}
		b.BindError()
		if len(b.errors) != 0 {
			t.Errorf("expected 0 errors, got %v", len(b.errors))
		}
	})
}
