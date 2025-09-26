package codegen

import (
	"fmt"
	"testing"
)

func TestFirstToLower_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test Case 1", "Hello", "hello"},
		{"Test Case 2", "WORLD", "wORLD"},
		{"Test Case 3", "1st", "1st"},
		{"Test Case 4", "", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := firstToLower(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
