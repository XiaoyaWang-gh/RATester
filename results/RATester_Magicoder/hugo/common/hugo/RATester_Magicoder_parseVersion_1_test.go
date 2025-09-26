package hugo

import (
	"fmt"
	"testing"
)

func TestparseVersion_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected [3]int
	}{
		{
			name:     "valid version",
			input:    "1.2.3",
			expected: [3]int{1, 2, 3},
		},
		{
			name:     "invalid version",
			input:    "1.2",
			expected: [3]int{1, 2, 0},
		},
		{
			name:     "empty version",
			input:    "",
			expected: [3]int{0, 0, 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			major, minor, patch := parseVersion(test.input)
			if major != test.expected[0] || minor != test.expected[1] || patch != test.expected[2] {
				t.Errorf("Expected %v, got %v, %v, %v", test.expected, major, minor, patch)
			}
		})
	}
}
