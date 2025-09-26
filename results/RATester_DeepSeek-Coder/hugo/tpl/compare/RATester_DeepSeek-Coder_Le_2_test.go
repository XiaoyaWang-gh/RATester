package compare

import (
	"fmt"
	"testing"
)

func TestLe_2(t *testing.T) {
	n := &Namespace{}

	t.Run("should return true if first argument is less than or equal to all others", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := n.Le(1, 1, 2, 3)
		if !result {
			t.Errorf("Expected true, got %v", result)
		}
	})

	t.Run("should return false if first argument is not less than or equal to all others", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := n.Le(4, 1, 2, 3)
		if result {
			t.Errorf("Expected false, got %v", result)
		}
	})

	t.Run("should return true if first argument is equal to all others", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := n.Le(1, 1, 1)
		if !result {
			t.Errorf("Expected true, got %v", result)
		}
	})

	t.Run("should return true if first argument is less than or equal to itself", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := n.Le(1, 1)
		if !result {
			t.Errorf("Expected true, got %v", result)
		}
	})
}
