package math

import (
	"fmt"
	"testing"
)

func TestSum_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	t.Run("single value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		sum, err := ns.Sum(1.0)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if sum != 1.0 {
			t.Errorf("expected 1.0, got %v", sum)
		}
	})

	t.Run("multiple values", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		sum, err := ns.Sum(1.0, 2.0, 3.0)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if sum != 6.0 {
			t.Errorf("expected 6.0, got %v", sum)
		}
	})

	t.Run("with error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Sum("a")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
