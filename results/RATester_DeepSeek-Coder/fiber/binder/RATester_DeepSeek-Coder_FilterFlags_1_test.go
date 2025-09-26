package binder

import (
	"fmt"
	"testing"
)

func TestFilterFlags_1(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected string
	}

	tests := []test{
		{
			name:     "Test case 1",
			input:    "Hello World",
			expected: "Hello",
		},
		{
			name:     "Test case 2",
			input:    "Hello;World",
			expected: "Hello",
		},
		{
			name:     "Test case 3",
			input:    "Hello World;",
			expected: "Hello",
		},
		{
			name:     "Test case 4",
			input:    "Hello;World;",
			expected: "Hello",
		},
		{
			name:     "Test case 5",
			input:    "Hello World;;;;",
			expected: "Hello",
		},
		{
			name:     "Test case 6",
			input:    "Hello;;;World;;;",
			expected: "Hello",
		},
		{
			name:     "Test case 7",
			input:    "Hello World",
			expected: "Hello",
		},
		{
			name:     "Test case 8",
			input:    "Hello;;;;World",
			expected: "Hello",
		},
		{
			name:     "Test case 9",
			input:    "Hello;;;;;World",
			expected: "Hello",
		},
		{
			name:     "Test case 10",
			input:    "Hello;;;;;;World",
			expected: "Hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := FilterFlags(tt.input)
			if got != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, got)
			}
		})
	}
}
