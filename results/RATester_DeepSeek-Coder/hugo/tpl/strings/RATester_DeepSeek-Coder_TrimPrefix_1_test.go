package strings

import (
	"fmt"
	"testing"
)

func TestTrimPrefix_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		prefix := "pre"
		s := "prefixstring"
		expected := "string"
		result, err := ns.TrimPrefix(prefix, s)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})

	t.Run("error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		prefix := 123
		s := "prefixstring"
		_, err := ns.TrimPrefix(prefix, s)
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
}
