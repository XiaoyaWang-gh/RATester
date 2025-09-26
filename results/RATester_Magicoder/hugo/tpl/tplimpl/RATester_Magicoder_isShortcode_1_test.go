package tplimpl

import (
	"fmt"
	"testing"
)

func TestisShortcode_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Shortcode with prefix",
			input:    "shortcodesPathPrefix/test",
			expected: true,
		},
		{
			name:     "Shortcode without prefix",
			input:    "test",
			expected: false,
		},
		{
			name:     "Empty string",
			input:    "",
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

			got := isShortcode(tt.input)
			if got != tt.expected {
				t.Errorf("isShortcode(%s) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
