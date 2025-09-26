package strings

import (
	"fmt"
	"testing"
)

func TestToUpper_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.ToUpper("test")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != "TEST" {
			t.Errorf("Expected 'TEST', got '%s'", result)
		}
	})

	t.Run("error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.ToUpper(123)
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
}
