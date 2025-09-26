package allconfig

import (
	"fmt"
	"testing"
)

func TestCreateTitle_1(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected string
	}

	testCases := []testCase{
		{
			name:     "Test with a simple string",
			input:    "test",
			expected: "test",
		},
		{
			name:     "Test with a string with spaces",
			input:    "test with spaces",
			expected: "test with spaces",
		},
		{
			name:     "Test with a string with special characters",
			input:    "test@with#special*characters",
			expected: "test@with#special*characters",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := ConfigLanguage{}
			result := c.CreateTitle(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
