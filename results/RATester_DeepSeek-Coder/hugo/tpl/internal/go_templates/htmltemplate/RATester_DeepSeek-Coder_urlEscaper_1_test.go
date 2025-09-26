package template

import (
	"fmt"
	"testing"
)

func TestUrlEscaper_1(t *testing.T) {
	t.Run("Test with string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := urlEscaper("test")
		expected := "test"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Test with int", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := urlEscaper(123)
		expected := "123"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Test with float", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := urlEscaper(123.456)
		expected := "123.456"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Test with bool", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := urlEscaper(true)
		expected := "true"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}
