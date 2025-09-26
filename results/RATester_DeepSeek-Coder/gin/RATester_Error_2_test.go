package gin

import (
	"errors"
	"fmt"
	"testing"
)

func TestError_2(t *testing.T) {
	t.Run("should return an error if the error is nil", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected a panic, but no panic occurred")
			}
		}()

		c := &Context{}
		c.Error(nil)
	})

	t.Run("should return an error if the error is not of type *Error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		c := &Context{}
		err := errors.New("test error")
		parsedError := c.Error(err)

		if parsedError.Err != err {
			t.Errorf("Expected error to be %v, but got %v", err, parsedError.Err)
		}

		if parsedError.Type != ErrorTypePrivate {
			t.Errorf("Expected error type to be %v, but got %v", ErrorTypePrivate, parsedError.Type)
		}
	})

	t.Run("should return the existing error if the error is of type *Error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		c := &Context{}
		existingError := &Error{
			Err:  errors.New("existing error"),
			Type: ErrorTypePrivate,
		}
		parsedError := c.Error(existingError)

		if parsedError.Err != existingError.Err {
			t.Errorf("Expected error to be %v, but got %v", existingError.Err, parsedError.Err)
		}

		if parsedError.Type != existingError.Type {
			t.Errorf("Expected error type to be %v, but got %v", existingError.Type, parsedError.Type)
		}
	})
}
