package math

import (
	"fmt"
	"testing"
)

func TestToRadians_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	t.Run("valid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.ToRadians(45)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 0.7853981633974483 {
			t.Errorf("Expected 0.7853981633974483, got %v", result)
		}
	})

	t.Run("invalid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.ToRadians("not a number")
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
}
