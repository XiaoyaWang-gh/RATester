package template

import (
	"fmt"
	"testing"
)

func TestRcdataEscaper_1(t *testing.T) {
	t.Run("Test with contentTypeHTML", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := rcdataEscaper("<script>alert('XSS')</script>", contentTypeHTML)
		expected := "&lt;script&gt;alert('XSS')&lt;/script&gt;"
		if result != expected {
			t.Errorf("Expected '%s', but got '%s'", expected, result)
		}
	})

	t.Run("Test with other contentType", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := rcdataEscaper("<script>alert('XSS')</script>", 1)
		expected := "&lt;script&gt;alert('XSS')&lt;/script&gt;"
		if result != expected {
			t.Errorf("Expected '%s', but got '%s'", expected, result)
		}
	})
}
