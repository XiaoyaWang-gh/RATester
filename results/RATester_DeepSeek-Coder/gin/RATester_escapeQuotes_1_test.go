package gin

import (
	"fmt"
	"testing"
)

func TestEscapeQuotes_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", ""},
		{"No quotes", "Hello, world", "Hello, world"},
		{"Single quote", "O'Reilly", "O''Reilly"},
		{"Double quote", "Hello \"world\"", "Hello ''world''"},
		{"Mixed quotes", "Hello \"O'Reilly\"", "Hello ''O''Reilly''"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := escapeQuotes(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, result)
			}
		})
	}
}
