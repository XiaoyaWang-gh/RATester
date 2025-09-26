package template

import (
	"fmt"
	"testing"
)

func TestdoublePercent_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test case 1",
			input:    "Hello%World",
			expected: "Hello%%World",
		},
		{
			name:     "Test case 2",
			input:    "%%%",
			expected: "%%%%%",
		},
		{
			name:     "Test case 3",
			input:    "",
			expected: "",
		},
		{
			name:     "Test case 4",
			input:    "No%%PercentSign",
			expected: "No%%PercentSign",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := doublePercent(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
