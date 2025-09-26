package task

import (
	"fmt"
	"testing"
)

func TestparseIntOrName_1(t *testing.T) {
	names := map[string]uint{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	tests := []struct {
		name     string
		expr     string
		expected uint
	}{
		{"Test One", "one", 1},
		{"Test Two", "two", 2},
		{"Test Three", "three", 3},
		{"Test Four", "four", 4},
		{"Test Five", "5", 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := parseIntOrName(test.expr, names)
			if result != test.expected {
				t.Errorf("Expected %d, but got %d", test.expected, result)
			}
		})
	}
}
