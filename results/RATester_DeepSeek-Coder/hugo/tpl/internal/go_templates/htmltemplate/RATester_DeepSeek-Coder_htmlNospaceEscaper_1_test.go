package template

import (
	"fmt"
	"testing"
)

func TestHtmlNospaceEscaper_1(t *testing.T) {
	t.Run("Test with empty string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := htmlNospaceEscaper("")
		if result != filterFailsafe {
			t.Errorf("Expected %s, got %s", filterFailsafe, result)
		}
	})

	t.Run("Test with contentTypeHTML", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := htmlNospaceEscaper("<p>Hello, World</p>", contentTypeHTML)
		expected := "Hello,World"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Test with regular string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := htmlNospaceEscaper("Hello, World")
		expected := "Hello,World"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}
