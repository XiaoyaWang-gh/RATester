package transform

import (
	"fmt"
	"testing"
)

func TestHTMLUnescape_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("valid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		input := "&lt;div&gt;Hello, World&lt;/div&gt;"
		expected := "<div>Hello, World</div>"

		output, err := ns.HTMLUnescape(input)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if output != expected {
			t.Errorf("Expected %q, got %q", expected, output)
		}
	})

	t.Run("invalid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		input := "<div>Hello, World</div>"
		expected := "<div>Hello, World</div>"

		output, err := ns.HTMLUnescape(input)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}

		if output != expected {
			t.Errorf("Expected %q, got %q", expected, output)
		}
	})
}
