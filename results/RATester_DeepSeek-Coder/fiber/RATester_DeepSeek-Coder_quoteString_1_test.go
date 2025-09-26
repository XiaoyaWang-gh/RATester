package fiber

import (
	"fmt"
	"testing"
)

func TestQuoteString_1(t *testing.T) {
	app := &App{
		getString: func(b []byte) string {
			return string(b)
		},
	}

	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Single word",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "Multiple words",
			input:    "hello world",
			expected: "hello world",
		},
		{
			name:     "Special characters",
			input:    "@$%^&*()",
			expected: "@$%^&*()",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := app.quoteString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
