package math

import (
	"fmt"
	"testing"
)

func TestPow_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Pow(2, 3)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 8 {
			t.Errorf("Expected 8, got %v", result)
		}
	})

	t.Run("error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Pow("2", 3)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
