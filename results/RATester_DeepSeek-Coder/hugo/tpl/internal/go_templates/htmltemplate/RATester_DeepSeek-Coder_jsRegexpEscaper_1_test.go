package template

import (
	"fmt"
	"testing"
)

func TestJsRegexpEscaper_1(t *testing.T) {
	t.Run("Test with empty input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := jsRegexpEscaper()
		expected := "(?:)"
		if result != expected {
			t.Errorf("Expected '%s', but got '%s'", expected, result)
		}
	})

	t.Run("Test with non-empty input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := jsRegexpEscaper("test")
		expected := "test"
		if result != expected {
			t.Errorf("Expected '%s', but got '%s'", expected, result)
		}
	})
}
