package compare

import (
	"fmt"
	"testing"
)

func TestConditional_1(t *testing.T) {
	n := &Namespace{}

	t.Run("TestConditional_True", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := "v1"
		result := n.Conditional(true, "v1", "v2")
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("TestConditional_False", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := "v2"
		result := n.Conditional(false, "v1", "v2")
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
