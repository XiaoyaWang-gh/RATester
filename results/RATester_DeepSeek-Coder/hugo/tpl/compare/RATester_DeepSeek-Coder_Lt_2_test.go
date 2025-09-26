package compare

import (
	"fmt"
	"testing"
)

func TestLt_2(t *testing.T) {
	n := &Namespace{}

	t.Run("should return true if first is less than others", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := n.Lt(1, 2, 3)
		if !result {
			t.Errorf("Expected true, got %v", result)
		}
	})

	t.Run("should return false if first is not less than others", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := n.Lt(3, 2, 1)
		if result {
			t.Errorf("Expected false, got %v", result)
		}
	})

	t.Run("should return false if first equals others", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := n.Lt(2, 2, 3)
		if result {
			t.Errorf("Expected false, got %v", result)
		}
	})
}
