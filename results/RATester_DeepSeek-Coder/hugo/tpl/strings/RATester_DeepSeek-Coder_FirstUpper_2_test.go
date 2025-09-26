package strings

import (
	"fmt"
	"testing"
)

func TestFirstUpper_2(t *testing.T) {
	ns := &Namespace{}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.FirstUpper("test")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != "Test" {
			t.Errorf("Expected 'Test', got '%s'", result)
		}
	})

	t.Run("error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.FirstUpper(123)
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
}
