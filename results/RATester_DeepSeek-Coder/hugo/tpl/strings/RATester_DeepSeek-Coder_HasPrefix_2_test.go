package strings

import (
	"fmt"
	"testing"
)

func TestHasPrefix_2(t *testing.T) {
	ns := &Namespace{}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got, err := ns.HasPrefix("hello world", "hello")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !got {
			t.Errorf("expected true, got false")
		}
	})

	t.Run("failure", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got, err := ns.HasPrefix("hello world", "world")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if got {
			t.Errorf("expected false, got true")
		}
	})

	t.Run("error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.HasPrefix(123, "hello")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
