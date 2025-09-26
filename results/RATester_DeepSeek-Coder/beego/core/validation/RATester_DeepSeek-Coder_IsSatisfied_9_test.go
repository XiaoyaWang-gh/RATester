package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_9(t *testing.T) {
	e := Enum{
		Rules: "admin|user",
		Key:   "user",
	}

	type test struct {
		name     string
		input    interface{}
		expected bool
	}

	tests := []test{
		{
			name:     "Test Case 1",
			input:    "admin",
			expected: true,
		},
		{
			name:     "Test Case 2",
			input:    "user",
			expected: true,
		},
		{
			name:     "Test Case 3",
			input:    "guest",
			expected: false,
		},
		{
			name:     "Test Case 4",
			input:    123,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := e.IsSatisfied(tt.input)
			if got != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, got)
			}
		})
	}
}
