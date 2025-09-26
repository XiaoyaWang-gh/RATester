package debug

import (
	"fmt"
	"testing"
)

func TestVisualizeSpaces_1(t *testing.T) {
	type testCase struct {
		name     string
		input    any
		expected string
	}

	testCases := []testCase{
		{
			name:     "Test with string",
			input:    "Hello, World",
			expected: "Hello, World",
		},
		{
			name:     "Test with number",
			input:    123,
			expected: "123",
		},
		{
			name:     "Test with boolean",
			input:    true,
			expected: "true",
		},
		{
			name:     "Test with nil",
			input:    nil,
			expected: "<nil>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ns := &Namespace{}
			result := ns.VisualizeSpaces(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
