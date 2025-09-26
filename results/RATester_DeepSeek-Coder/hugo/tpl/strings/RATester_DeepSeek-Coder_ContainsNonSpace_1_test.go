package strings

import (
	"fmt"
	"testing"
)

func TestContainsNonSpace_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("contains non-space characters", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got, err := ns.ContainsNonSpace(" Hello World ")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !got {
			t.Errorf("expected true, got false")
		}
	})

	t.Run("does not contain non-space characters", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got, err := ns.ContainsNonSpace("     ")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if got {
			t.Errorf("expected false, got true")
		}
	})

	t.Run("invalid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.ContainsNonSpace(123)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
