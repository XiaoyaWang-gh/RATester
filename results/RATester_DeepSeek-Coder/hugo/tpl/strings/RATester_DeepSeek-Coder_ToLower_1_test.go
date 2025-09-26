package strings

import (
	"fmt"
	"testing"
)

func TestToLower_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.ToLower("HELLO")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != "hello" {
			t.Errorf("Expected 'hello', got '%s'", result)
		}
	})

	t.Run("error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.ToLower(123)
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
}
