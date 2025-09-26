package strings

import (
	"fmt"
	"testing"
)

func TestContains_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("contains true", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got, err := ns.Contains("hello world", "world")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !got {
			t.Errorf("expected true, got false")
		}
	})

	t.Run("contains false", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got, err := ns.Contains("hello world", "foo")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if got {
			t.Errorf("expected false, got true")
		}
	})

	t.Run("contains error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Contains(123, "foo")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
