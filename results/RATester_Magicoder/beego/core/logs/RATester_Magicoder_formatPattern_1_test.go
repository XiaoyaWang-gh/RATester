package logs

import (
	"fmt"
	"testing"
)

func TestformatPattern_1(t *testing.T) {
	type test struct {
		name     string
		f        interface{}
		v        []interface{}
		expected string
	}

	tests := []test{
		{
			name:     "Test case 1",
			f:        "Hello, %s!",
			v:        []interface{}{"World"},
			expected: "Hello, World!",
		},
		{
			name:     "Test case 2",
			f:        "Number: %d",
			v:        []interface{}{42},
			expected: "Number: 42",
		},
		{
			name:     "Test case 3",
			f:        "No format specifier",
			v:        []interface{}{},
			expected: "No format specifier",
		},
		{
			name:     "Test case 4",
			f:        "Multiple format specifiers: %s, %d",
			v:        []interface{}{"World", 42},
			expected: "Multiple format specifiers: World, 42",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := formatPattern(tt.f, tt.v...)
			if got != tt.expected {
				t.Errorf("formatPattern(%v, %v) = %v, want %v", tt.f, tt.v, got, tt.expected)
			}
		})
	}
}
