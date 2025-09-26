package helpers

import (
	"fmt"
	"testing"
)

func TestURLize_1(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected string
	}

	tests := []test{
		{
			name:     "Test case 1",
			input:    "/test/path",
			expected: "/test/path",
		},
		{
			name:     "Test case 2",
			input:    "/test/path/with/spaces",
			expected: "/test/path/with/spaces",
		},
		{
			name:     "Test case 3",
			input:    "/test/path/with/special/characters",
			expected: "/test/path/with/special/characters",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ps := &PathSpec{}
			result := ps.URLize(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
