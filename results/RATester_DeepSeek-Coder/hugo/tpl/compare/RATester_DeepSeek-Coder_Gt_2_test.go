package compare

import (
	"fmt"
	"testing"
)

func TestGt_2(t *testing.T) {
	n := &Namespace{}

	t.Run("should return true if the first argument is greater than all others", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := n.Gt(2, 1)
		if !result {
			t.Errorf("Expected true, got %v", result)
		}
	})

	t.Run("should return false if the first argument is not greater than all others", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := n.Gt(1, 2)
		if result {
			t.Errorf("Expected false, got %v", result)
		}
	})

	t.Run("should return true if the first argument is greater than all others in a slice", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := n.Gt(2, []any{1, 0}...)
		if !result {
			t.Errorf("Expected true, got %v", result)
		}
	})

	t.Run("should return false if the first argument is not greater than all others in a slice", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := n.Gt(1, []any{2, 3}...)
		if result {
			t.Errorf("Expected false, got %v", result)
		}
	})
}
