package template

import (
	"fmt"
	"testing"
)

func TestContainsSpecialScriptTag_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected bool
	}{
		{
			name:     "contains special script tag",
			input:    []byte("<script>alert('xss')</script>"),
			expected: true,
		},
		{
			name:     "does not contain special script tag",
			input:    []byte("<p>This is a paragraph.</p>"),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := containsSpecialScriptTag(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
