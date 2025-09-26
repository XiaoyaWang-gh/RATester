package math

import (
	"fmt"
	"testing"
)

func TestCos_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	t.Run("valid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Cos(30)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		expected := 0.8660254037844386
		if result != expected {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("invalid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Cos("invalid")
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	})
}
