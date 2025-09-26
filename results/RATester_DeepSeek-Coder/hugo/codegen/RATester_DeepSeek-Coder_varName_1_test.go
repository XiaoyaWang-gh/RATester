package codegen

import (
	"fmt"
	"testing"
)

func TestVarName_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test case 1", "type", "typ"},
		{"Test case 2", "package", "pkg"},
		{"Test case 3", "len", "length"},
		{"Test case 4", "var", "var"},
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
				t.Errorf("Expected %s, got %s", test.expected, result)
			}
		})
	}
}
