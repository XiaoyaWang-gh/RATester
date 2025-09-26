package math

import (
	"fmt"
	"math"
	"testing"
)

func TestToDegrees_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.ToDegrees(math.Pi)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result != 180 {
			t.Errorf("expected 180, got %v", result)
		}
	})

	t.Run("error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.ToDegrees("not a number")
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})
}
