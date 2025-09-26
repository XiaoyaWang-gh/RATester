package hstrings

import (
	"fmt"
	"testing"
)

func TestEqualFold_1(t *testing.T) {
	tests := []struct {
		name     string
		s        StringEqualFold
		s2       string
		expected bool
	}{
		{
			name:     "EqualFold_True",
			s:        "Hello",
			s2:       "HELLO",
			expected: true,
		},
		{
			name:     "EqualFold_False",
			s:        "Hello",
			s2:       "WORLD",
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

			result := tt.s.EqualFold(tt.s2)
			if result != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
