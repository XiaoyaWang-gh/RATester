package math

import (
	"fmt"
	"math"
	"testing"
)

func TestAcos_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("valid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Acos(1.0)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if math.Abs(result-math.Acos(1.0)) > 1e-9 {
			t.Errorf("expected %v, got %v", math.Acos(1.0), result)
		}
	})

	t.Run("invalid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Acos("invalid")
		if err == nil {
			t.Errorf("expected an error")
		}
	})
}
