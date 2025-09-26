package lang

import (
	"fmt"
	"testing"
)

func TestcastPrecisionNumber_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("valid precision and number", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		p, n, err := ns.castPrecisionNumber(10, 123.456)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if p != 10 {
			t.Errorf("expected precision 10, got %d", p)
		}
		if n != 123.456 {
			t.Errorf("expected number 123.456, got %f", n)
		}
	})

	t.Run("invalid precision", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, _, err := ns.castPrecisionNumber("invalid", 123.456)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("invalid number", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, _, err := ns.castPrecisionNumber(10, "invalid")
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("precision too high", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, _, err := ns.castPrecisionNumber(21, 123.456)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
