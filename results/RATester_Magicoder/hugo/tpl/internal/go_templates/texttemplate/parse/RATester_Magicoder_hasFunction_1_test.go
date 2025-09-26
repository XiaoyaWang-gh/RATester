package parse

import (
	"fmt"
	"testing"
)

func TesthasFunction_1(t *testing.T) {
	tree := &Tree{
		funcs: []map[string]any{
			{"func1": nil, "func2": nil},
			{"func3": nil, "func4": nil},
		},
	}

	tests := []struct {
		name     string
		funcName string
		expected bool
	}{
		{"Test case 1", "func1", true},
		{"Test case 2", "func5", false},
		{"Test case 3", "func3", true},
		{"Test case 4", "func4", true},
		{"Test case 5", "func6", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tree.hasFunction(test.funcName)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
