package page

import (
	"fmt"
	"testing"
)

func TestConcatLast_1(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected string
	}

	tests := []test{
		{
			name:     "ConcatLast_Empty",
			input:    "",
			expected: "",
		},
		{
			name:     "ConcatLast_Single",
			input:    "test",
			expected: "test",
		},
		{
			name:     "ConcatLast_Multiple",
			input:    "test/test",
			expected: "test/test",
		},
		{
			name:     "ConcatLast_TrailingSlash",
			input:    "test/",
			expected: "test",
		},
		{
			name:     "ConcatLast_TrailingSlash_Multiple",
			input:    "test/test/",
			expected: "test/test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &pagePathBuilder{}
			p.ConcatLast(tt.input)
			if p.Last() != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, p.Last())
			}
		})
	}
}
