package strings

import (
	"fmt"
	"testing"
)

func TestTrimRight_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name     string
		cutset   any
		s        any
		expected string
	}{
		{
			name:     "TrimRight with string",
			cutset:   "abc",
			s:        "hello worldabc",
			expected: "hello world",
		},
		{
			name:     "TrimRight with rune",
			cutset:   'a',
			s:        "hello worldaaa",
			expected: "hello world",
		},
		{
			name:     "TrimRight with empty string",
			cutset:   "",
			s:        "hello world",
			expected: "hello world",
		},
		{
			name:     "TrimRight with non-existing string",
			cutset:   "xyz",
			s:        "hello world",
			expected: "hello world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := ns.TrimRight(tt.cutset, tt.s)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}
