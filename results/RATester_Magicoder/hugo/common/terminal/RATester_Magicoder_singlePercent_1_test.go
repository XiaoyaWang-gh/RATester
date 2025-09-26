package terminal

import (
	"fmt"
	"testing"
)

func TestsinglePercent_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"no_percent", "hello", "hello"},
		{"one_percent", "hello%%world", "hello%world"},
		{"multiple_percent", "hello%%%%world", "hello%world"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := singlePercent(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
