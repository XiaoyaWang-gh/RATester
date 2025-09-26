package math

import (
	"fmt"
	"testing"
)

func TestRound_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Round(1.5)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 2 {
			t.Errorf("Expected 2, got %v", result)
		}
	})

	t.Run("error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Round("not a number")
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
}
