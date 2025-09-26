package fiber

import (
	"fmt"
	"testing"
)

func TestfindNextParamPosition_1(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		expected int
	}{
		{
			name:     "no parameters",
			pattern:  "no parameters",
			expected: -1,
		},
		{
			name:     "one parameter",
			pattern:  "one parameter:%s",
			expected: 13,
		},
		{
			name:     "two parameters",
			pattern:  "two parameters:%s and %s",
			expected: 17,
		},
		{
			name:     "three parameters",
			pattern:  "three parameters:%s, %s, and %s",
			expected: 22,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := findNextParamPosition(test.pattern)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}
