package template

import (
	"fmt"
	"testing"
)

func TestEvalArgs_2(t *testing.T) {
	t.Run("Test with single string argument", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := evalArgs([]any{"Hello, World!"})
		expected := "Hello, World!"
		if result != expected {
			t.Errorf("Expected '%s', got '%s'", expected, result)
		}
	})

	t.Run("Test with multiple arguments", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := evalArgs([]any{1, 2, 3, "Hello, World!"})
		expected := "1 2 3 Hello, World!"
		if result != expected {
			t.Errorf("Expected '%s', got '%s'", expected, result)
		}
	})

	t.Run("Test with unprintable arguments", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		type unprintable struct{}
		result := evalArgs([]any{unprintable{}, 1, 2, 3, "Hello, World!"})
		expected := "[unprintable] 1 2 3 Hello, World!"
		if result != expected {
			t.Errorf("Expected '%s', got '%s'", expected, result)
		}
	})
}
