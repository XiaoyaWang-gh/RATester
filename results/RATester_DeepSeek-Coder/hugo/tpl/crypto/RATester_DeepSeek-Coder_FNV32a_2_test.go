package crypto

import (
	"fmt"
	"testing"
)

func TestFNV32a_2(t *testing.T) {
	ns := &Namespace{}

	t.Run("valid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.FNV32a("test")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result == 0 {
			t.Errorf("Expected non-zero result, got 0")
		}
	})

	t.Run("invalid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.FNV32a(123)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
