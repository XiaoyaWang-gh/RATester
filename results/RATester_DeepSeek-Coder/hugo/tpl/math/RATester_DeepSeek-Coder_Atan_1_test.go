package math

import (
	"fmt"
	"math"
	"testing"
)

func TestAtan_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	t.Run("valid", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Atan(1.0)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != math.Atan(1.0) {
			t.Errorf("Expected %v, got %v", math.Atan(1.0), result)
		}
	})

	t.Run("invalid", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Atan("invalid")
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
}
