package compare

import (
	"fmt"
	"testing"
)

func TestCheckComparisonArgCount_1(t *testing.T) {
	n := &Namespace{}

	t.Run("Test with minimum arguments", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := n.checkComparisonArgCount(1, 1)
		if result != true {
			t.Errorf("Expected true, got %v", result)
		}
	})

	t.Run("Test with insufficient arguments", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected a panic, but no panic occurred")
			}
		}()
		n.checkComparisonArgCount(2)
	})
}
