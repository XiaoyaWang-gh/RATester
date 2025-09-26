package gin

import (
	"fmt"
	"testing"
)

func TestLastChar_1(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected uint8
	}

	tests := []test{
		{"Test Case 1", "Hello", 'o'},
		{"Test Case 2", "World", 'd'},
		{"Test Case 3", "Golang", 'g'},
		{"Test Case 4", "", '0'},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := lastChar(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
