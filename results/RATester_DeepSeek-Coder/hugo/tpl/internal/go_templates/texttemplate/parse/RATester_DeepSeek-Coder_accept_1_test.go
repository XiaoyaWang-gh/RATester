package parse

import (
	"fmt"
	"testing"
)

func TestAccept_1(t *testing.T) {
	type test struct {
		name     string
		input    string
		valid    string
		expected bool
	}

	tests := []test{
		{
			name:     "valid input",
			input:    "abc",
			valid:    "abc",
			expected: true,
		},
		{
			name:     "invalid input",
			input:    "def",
			valid:    "abc",
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

			l := &lexer{
				input: tt.input,
			}
			result := l.accept(tt.valid)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
