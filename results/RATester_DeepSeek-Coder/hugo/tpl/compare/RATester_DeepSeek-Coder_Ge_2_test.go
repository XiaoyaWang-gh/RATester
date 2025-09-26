package compare

import (
	"fmt"
	"testing"
)

func TestGe_2(t *testing.T) {
	n := &Namespace{}

	t.Run("should return true if first is greater than or equal to all others", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := n.Ge(10, 5, 10)
		if !result {
			t.Errorf("Expected true, got %v", result)
		}
	})

	t.Run("should return false if first is not greater than or equal to any others", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := n.Ge(10, 15, 20)
		if result {
			t.Errorf("Expected false, got %v", result)
		}
	})

	t.Run("should return true if first is equal to others", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := n.Ge(10, 10, 10)
		if !result {
			t.Errorf("Expected true, got %v", result)
		}
	})

	t.Run("should return true if first is greater than others", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := n.Ge(10, 5, 8)
		if !result {
			t.Errorf("Expected true, got %v", result)
		}
	})
}
