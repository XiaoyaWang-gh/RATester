package template

import (
	"fmt"
	"testing"
)

func TestUrlNormalizer_1(t *testing.T) {
	t.Run("Test with valid URL", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := urlNormalizer("https://www.example.com")
		expected := "https://www.example.com"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Test with invalid URL", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := urlNormalizer("www.example.com")
		expected := "http://www.example.com"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Test with empty URL", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := urlNormalizer("")
		expected := ""
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}
