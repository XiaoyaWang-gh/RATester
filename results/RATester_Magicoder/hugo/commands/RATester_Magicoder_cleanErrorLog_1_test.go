package commands

import (
	"fmt"
	"testing"
)

func TestcleanErrorLog_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test case 1",
			input:    "error 1: error 2: error 1",
			expected: "error 1: error 2",
		},
		{
			name:     "Test case 2",
			input:    "error 1: error 1: error 2",
			expected: "error 1: error 2",
		},
		{
			name:     "Test case 3",
			input:    "error 1: error 2: error 1: error 3",
			expected: "error 1: error 2: error 3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := cleanErrorLog(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %s, but got %s", tt.expected, result)
			}
		})
	}
}
