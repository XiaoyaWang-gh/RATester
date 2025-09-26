package herrors

import (
	"errors"
	"fmt"
	"testing"
)

func TestUnwrapFileError_1(t *testing.T) {
	t.Run("should return nil if no FileError is found", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		err := errors.New("some error")
		result := UnwrapFileError(err)
		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})

	t.Run("should return the innermost FileError if it exists", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		innerErr := &fileError{fileType: "inner"}
		outerErr := fmt.Errorf("outer: %w", innerErr)
		result := UnwrapFileError(outerErr)
		if result != innerErr {
			t.Errorf("Expected %v, got %v", innerErr, result)
		}
	})

	t.Run("should return the innermost FileError if it exists in a chain", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		innerErr := &fileError{fileType: "inner"}
		midErr := fmt.Errorf("mid: %w", innerErr)
		outerErr := fmt.Errorf("outer: %w", midErr)
		result := UnwrapFileError(outerErr)
		if result != innerErr {
			t.Errorf("Expected %v, got %v", innerErr, result)
		}
	})
}
