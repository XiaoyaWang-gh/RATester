package math

import (
	"fmt"
	"math"
	"testing"
)

func TestMax_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	t.Run("single input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		max, err := ns.Max(1.0)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if max != 1.0 {
			t.Errorf("expected 1.0, got %v", max)
		}
	})

	t.Run("multiple inputs", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		max, err := ns.Max(1.0, 2.0, 3.0)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if max != 3.0 {
			t.Errorf("expected 3.0, got %v", max)
		}
	})

	t.Run("with NaN", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Max(math.NaN())
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
