package template

import (
	"fmt"
	"testing"
)

func TestIsTrue_1(t *testing.T) {
	t.Run("Test with true value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		truth, ok := IsTrue(true)
		if !truth || !ok {
			t.Errorf("Expected (true, true), got (%v, %v)", truth, ok)
		}
	})

	t.Run("Test with false value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		truth, ok := IsTrue(false)
		if truth || !ok {
			t.Errorf("Expected (false, true), got (%v, %v)", truth, ok)
		}
	})

	t.Run("Test with non-boolean value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		truth, ok := IsTrue("not a boolean")
		if truth || ok {
			t.Errorf("Expected (false, false), got (%v, %v)", truth, ok)
		}
	})
}
