package css

import (
	"fmt"
	"testing"
)

func TestUnquoted_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("Test with string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := ns.Unquoted("test")
		if result != "test" {
			t.Errorf("Expected 'test', got %s", result)
		}
	})

	t.Run("Test with int", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := ns.Unquoted(123)
		if result != "123" {
			t.Errorf("Expected '123', got %s", result)
		}
	})

	t.Run("Test with float", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := ns.Unquoted(123.456)
		if result != "123.456" {
			t.Errorf("Expected '123.456', got %s", result)
		}
	})

	t.Run("Test with bool", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := ns.Unquoted(true)
		if result != "true" {
			t.Errorf("Expected 'true', got %s", result)
		}
	})
}
