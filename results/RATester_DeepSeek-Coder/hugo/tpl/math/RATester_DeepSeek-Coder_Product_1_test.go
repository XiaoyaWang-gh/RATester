package math

import (
	"fmt"
	"testing"
)

func TestProduct_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	t.Run("single input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		product, err := ns.Product(2.0)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if product != 2.0 {
			t.Errorf("expected 2.0, got %v", product)
		}
	})

	t.Run("multiple inputs", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		product, err := ns.Product(2.0, 3.0, 4.0)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if product != 24.0 {
			t.Errorf("expected 24.0, got %v", product)
		}
	})

	t.Run("with non-numeric input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Product("a", 2.0)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("with empty input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Product()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}
