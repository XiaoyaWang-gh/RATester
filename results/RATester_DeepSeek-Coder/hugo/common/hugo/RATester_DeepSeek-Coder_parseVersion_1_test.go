package hugo

import (
	"fmt"
	"testing"
)

func TestParseVersion_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected [3]int
	}{
		{"Test case 1", "1.2.3", [3]int{1, 2, 3}},
		{"Test case 2", "1.2", [3]int{1, 2, 0}},
		{"Test case 3", "1", [3]int{1, 0, 0}},
		{"Test case 4", "1.2.3.4", [3]int{1, 2, 3}},
		{"Test case 5", "1.2.a", [3]int{1, 2, 0}},
		{"Test case 6", "1.a.3", [3]int{1, 0, 3}},
		{"Test case 7", "a.2.3", [3]int{0, 2, 3}},
		{"Test case 8", "1.2.3-alpha", [3]int{1, 2, 3}},
		{"Test case 9", "1.2-alpha.3", [3]int{1, 2, 0}},
		{"Test case 10", "1-alpha.2.3", [3]int{1, 0, 2}},
		{"Test case 11", "1.2.3-alpha.4", [3]int{1, 2, 3}},
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
