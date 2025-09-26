package codegen

import (
	"fmt"
	"testing"
)

func TestvarName_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test case 1", "type", "typ"},
		{"Test case 2", "package", "pkg"},
		{"Test case 3", "len", "length"},
		{"Test case 4", "name", "name"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := varName(test.input)
			if result != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
		})
	}
}
