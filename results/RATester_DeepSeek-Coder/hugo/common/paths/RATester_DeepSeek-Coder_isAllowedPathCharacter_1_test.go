package paths

import (
	"fmt"
	"testing"
)

func TestIsAllowedPathCharacter_1(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		i        int
		r        rune
		expected bool
	}{
		{
			name:     "allowed character",
			s:        "test",
			i:        0,
			r:        't',
			expected: true,
		},
		{
			name:     "disallowed character",
			s:        "test",
			i:        0,
			r:        ' ',
			expected: false,
		},
		{
			name:     "allowed special character",
			s:        "test",
			i:        0,
			r:        '.',
			expected: true,
		},
		{
			name:     "allowed percent sign",
			s:        "test%20",
			i:        4,
			r:        '%',
			expected: true,
		},
		{
			name:     "disallowed percent sign",
			s:        "test%20",
			i:        0,
			r:        '%',
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

			result := isAllowedPathCharacter(tc.s, tc.i, tc.r)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
