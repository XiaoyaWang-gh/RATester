package transform

import (
	"fmt"
	"html/template"
	"testing"
)

func TestEmojify_2(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	t.Run("valid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		input := ":smile:"
		expected := template.HTML("😄")

		result, err := ns.Emojify(input)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("invalid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		input := 123
		_, err := ns.Emojify(input)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
