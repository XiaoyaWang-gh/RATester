package transform

import (
	"fmt"
	"testing"
)

func TestStringToRune_1(t *testing.T) {
	t.Run("Test with valid string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		r, err := stringToRune("a")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if r != 'a' {
			t.Errorf("Expected 'a', got %v", r)
		}
	})

	t.Run("Test with empty string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		r, err := stringToRune("")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if r != 0 {
			t.Errorf("Expected 0, got %v", r)
		}
	})

	t.Run("Test with invalid string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := stringToRune("ab")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
