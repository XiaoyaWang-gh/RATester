package echo

import (
	"errors"
	"fmt"
	"testing"
)

func TestBindErrors_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		errors: []error{errors.New("error1"), errors.New("error2")},
	}

	errors := b.BindErrors()

	if len(errors) != 2 {
		t.Errorf("Expected 2 errors, got %d", len(errors))
	}

	if len(b.errors) != 0 {
		t.Errorf("Expected errors to be cleared, got %d", len(b.errors))
	}
}
