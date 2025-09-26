package constraints

import (
	"fmt"
	"testing"
)

func TestNotTagFunc_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected bool
	}{
		{
			name:     "Test case 1",
			input:    []string{"tag1", "tag2"},
			expected: false,
		},
		{
			name:     "Test case 2",
			input:    []string{"tag3", "tag4"},
			expected: true,
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

			result := notTagFunc(func(tags []string) bool {
				return len(tags) > 0
			})(tc.input)

			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
