package gin

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrors_1(t *testing.T) {
	t.Run("should return nil when errorMsgs is empty", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		errorMsgs := errorMsgs{}
		result := errorMsgs.Errors()
		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})

	t.Run("should return error strings when errorMsgs is not empty", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		errorMsgs := errorMsgs{&Error{Err: errors.New("error 1")}, &Error{Err: errors.New("error 2")}}
		result := errorMsgs.Errors()
		if len(result) != 2 {
			t.Errorf("Expected 2 errors, got %d", len(result))
		}
		if result[0] != "error 1" {
			t.Errorf("Expected 'error 1', got '%s'", result[0])
		}
		if result[1] != "error 2" {
			t.Errorf("Expected 'error 2', got '%s'", result[1])
		}
	})
}
