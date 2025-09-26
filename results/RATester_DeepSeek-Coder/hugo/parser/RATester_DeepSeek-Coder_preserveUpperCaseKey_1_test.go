package parser

import (
	"fmt"
	"testing"
)

func TestPreserveUpperCaseKey_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected bool
	}{
		{
			name:     "Test case 1",
			input:    []byte("Hello"),
			expected: true,
		},
		{
			name:     "Test case 2",
			input:    []byte("hello"),
			expected: false,
		},
		{
			name:     "Test case 3",
			input:    []byte("HELLO"),
			expected: true,
		},
		{
			name:     "Test case 4",
			input:    []byte("HELLO_WORLD"),
			expected: true,
		},
		{
			name:     "Test case 5",
			input:    []byte("_HELLO_WORLD"),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := preserveUpperCaseKey(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
