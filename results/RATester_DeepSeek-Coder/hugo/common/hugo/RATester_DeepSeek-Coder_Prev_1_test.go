package hugo

import (
	"fmt"
	"testing"
)

func TestPrev_1(t *testing.T) {
	type test struct {
		name     string
		input    Version
		expected Version
	}

	tests := []test{
		{
			name:     "TestPrev_1",
			input:    Version{Major: 0, Minor: 0},
			expected: Version{Major: 0, Minor: -1},
		},
		{
			name:     "TestPrev_2",
			input:    Version{Major: 1, Minor: 2},
			expected: Version{Major: 1, Minor: 1},
		},
		{
			name:     "TestPrev_3",
			input:    Version{Major: 2, Minor: 0},
			expected: Version{Major: 2, Minor: -1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.input.Prev()
			if result.Major != tc.expected.Major || result.Minor != tc.expected.Minor {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
