package math

import (
	"fmt"
	"math"
	"testing"
)

func TestMin_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	t.Run("single value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		min, err := ns.Min(1.0)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if min != 1.0 {
			t.Errorf("expected 1.0, got %v", min)
		}
	})

	t.Run("multiple values", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		min, err := ns.Min(1.0, 2.0, 3.0)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if min != 1.0 {
			t.Errorf("expected 1.0, got %v", min)
		}
	})

	t.Run("negative values", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		min, err := ns.Min(-1.0, -2.0, -3.0)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if min != -3.0 {
			t.Errorf("expected -3.0, got %v", min)
		}
	})

	t.Run("zero value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		min, err := ns.Min(0.0, 1.0, -1.0)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if min != -1.0 {
			t.Errorf("expected -1.0, got %v", min)
		}
	})

	t.Run("NaN value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Min(math.NaN())
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
