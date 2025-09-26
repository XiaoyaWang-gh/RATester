package compare

import (
	"fmt"
	"testing"
)

func TestNe_2(t *testing.T) {
	n := &Namespace{}

	t.Run("should return true if first and any other argument are not equal", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := n.Ne("test", "not equal")
		if !result {
			t.Errorf("Expected true, got %v", result)
		}
	})

	t.Run("should return false if first and any other argument are equal", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := n.Ne("test", "test")
		if result {
			t.Errorf("Expected false, got %v", result)
		}
	})

	t.Run("should return true if first and all other arguments are not equal", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := n.Ne("test", "not equal", "also not equal")
		if !result {
			t.Errorf("Expected true, got %v", result)
		}
	})

	t.Run("should return false if first and any of the other arguments are equal", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := n.Ne("test", "test", "also not equal")
		if result {
			t.Errorf("Expected false, got %v", result)
		}
	})
}
