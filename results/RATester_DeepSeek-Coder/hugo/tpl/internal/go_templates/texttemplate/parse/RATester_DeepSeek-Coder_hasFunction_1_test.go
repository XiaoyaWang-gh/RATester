package parse

import (
	"fmt"
	"testing"
)

func TestHasFunction_1(t *testing.T) {
	tree := &Tree{
		funcs: []map[string]any{
			{"func1": nil, "func2": nil},
			{"func3": nil, "func4": nil},
		},
	}

	tests := []struct {
		name     string
		funcName string
		want     bool
	}{
		{"Test Case 1", "func1", true},
		{"Test Case 2", "func5", false},
		{"Test Case 3", "func3", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tree.hasFunction(tt.funcName); got != tt.want {
				t.Errorf("hasFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
