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
			input:    []string{"tag1", "tag2", "tag3"},
			expected: false,
		},
		{
			name:     "Test case 2",
			input:    []string{"tag4", "tag5", "tag6"},
			expected: true,
		},
		{
			name:     "Test case 3",
			input:    []string{"tag7", "tag8", "tag9"},
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

			fn := notTagFunc(func(tags []string) bool {
				return len(tags) > 3
			})

			result := fn(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
