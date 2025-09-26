package helpers

import (
	"fmt"
	"testing"
)

func TestHasStringsPrefix_1(t *testing.T) {
	tests := []struct {
		name     string
		s        []string
		prefix   []string
		expected bool
	}{
		{
			name:     "Test case 1",
			s:        []string{"hello", "world"},
			prefix:   []string{"hello", "world"},
			expected: true,
		},
		{
			name:     "Test case 2",
			s:        []string{"hello", "world"},
			prefix:   []string{"hello"},
			expected: true,
		},
		{
			name:     "Test case 3",
			s:        []string{"hello", "world"},
			prefix:   []string{"world"},
			expected: false,
		},
		{
			name:     "Test case 4",
			s:        []string{"hello", "world"},
			prefix:   []string{"hello", "world", "again"},
			expected: false,
		},
		{
			name:     "Test case 5",
			s:        []string{"hello", "world"},
			prefix:   []string{},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := HasStringsPrefix(test.s, test.prefix)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
