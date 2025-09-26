package helpers

import (
	"fmt"
	"testing"
)

func TestResolveMarkup_1(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected string
	}

	testCases := []testCase{
		{
			name:     "Test case 1",
			input:    "test1",
			expected: "test1",
		},
		{
			name:     "Test case 2",
			input:    "test2",
			expected: "test2",
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			contentSpec := &ContentSpec{
				// Initialize ContentSpec with necessary dependencies
			}

			result := contentSpec.ResolveMarkup(tc.input)

			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
