package template

import (
	"fmt"
	"testing"
)

func TestTJS_1(t *testing.T) {
	type test struct {
		name     string
		input    []byte
		expected context
		length   int
	}

	tests := []test{
		{
			name:  "Test case 1",
			input: []byte(`"Hello, world!"`),
			expected: context{
				state: stateJSDqStr,
			},
			length: 14,
		},
		{
			name:  "Test case 2",
			input: []byte(`'Hello, world!'`),
			expected: context{
				state: stateJSSqStr,
			},
			length: 14,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual, length := tJS(context{}, tt.input)
			if actual.state != tt.expected.state {
				t.Errorf("Expected state %v, but got %v", tt.expected.state, actual.state)
			}
			if length != tt.length {
				t.Errorf("Expected length %v, but got %v", tt.length, length)
			}
		})
	}
}
