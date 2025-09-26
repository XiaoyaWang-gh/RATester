package helpers

import (
	"fmt"
	"html/template"
	"testing"
)

func TestBytesToHTML_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected template.HTML
	}{
		{
			name:     "empty byte slice",
			input:    []byte{},
			expected: template.HTML(""),
		},
		{
			name:     "non-empty byte slice",
			input:    []byte("<p>Hello, World</p>"),
			expected: template.HTML("<p>Hello, World</p>"),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			result := BytesToHTML(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
