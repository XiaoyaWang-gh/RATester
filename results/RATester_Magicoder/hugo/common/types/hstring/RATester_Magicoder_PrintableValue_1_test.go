package hstring

import (
	"fmt"
	"html/template"
	"testing"
)

func TestPrintableValue_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    HTML
		expected any
	}{
		{
			name:     "empty string",
			input:    "",
			expected: template.HTML(""),
		},
		{
			name:     "non-empty string",
			input:    "<p>Hello, World!</p>",
			expected: template.HTML("<p>Hello, World!</p>"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.input.PrintableValue()
			if result != tc.expected {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
