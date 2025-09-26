package web

import (
	"fmt"
	"testing"
)

func TestHtmlquote_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"simple string", "Hello, World", "Hello, World"},
		{"string with quotes", "“Hello”, World", "&ldquo;Hello&rdquo;, World"},
		{"string with space", "Hello World", "Hello&nbsp;World"},
		{"string with multiple spaces", "Hello   World", "Hello&nbsp;&nbsp;&nbsp;World"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := Htmlquote(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, result)
			}
		})
	}
}
