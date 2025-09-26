package strings

import (
	"fmt"
	"testing"
)

func TestCountRunes_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		input := "Hello, World!"
		expected := 13
		actual, err := ns.CountRunes(input)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		input := 123
		_, err := ns.CountRunes(input)
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
}
