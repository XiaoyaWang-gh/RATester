package math

import (
	"fmt"
	"testing"
)

func TestSqrt_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	t.Run("valid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Sqrt(4)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 2 {
			t.Errorf("Expected 2, got %v", result)
		}
	})

	t.Run("negative input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Sqrt(-4)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("non-numeric input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Sqrt("four")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
