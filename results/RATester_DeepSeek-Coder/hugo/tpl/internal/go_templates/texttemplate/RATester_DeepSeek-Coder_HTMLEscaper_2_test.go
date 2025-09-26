package template

import (
	"fmt"
	"testing"
)

func TestHTMLEscaper_2(t *testing.T) {
	t.Run("Test with string input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := "&lt;p&gt;Hello, World&lt;/p&gt;"
		result := HTMLEscaper("<p>Hello, World</p>")
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Test with int input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := "123"
		result := HTMLEscaper(123)
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Test with float input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := "123.456"
		result := HTMLEscaper(123.456)
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Test with bool input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := "true"
		result := HTMLEscaper(true)
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Test with nil input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := "<nil>"
		result := HTMLEscaper(nil)
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}
