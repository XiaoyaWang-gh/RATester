package math

import (
	"fmt"
	"math"
	"testing"
)

func TestSin_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	t.Run("valid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Sin(30)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if math.Abs(result-0.5) > 0.0001 {
			t.Errorf("Expected 0.5, got %v", result)
		}
	})

	t.Run("invalid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Sin("invalid")
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
}
