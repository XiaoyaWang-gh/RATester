package template

import (
	"fmt"
	"testing"
)

func TestUrlProcessor_1(t *testing.T) {
	type test struct {
		name     string
		norm     bool
		args     []any
		expected string
	}

	tests := []test{
		{
			name:     "Test with URL",
			norm:     false,
			args:     []any{"https://example.com"},
			expected: "https://example.com",
		},
		{
			name:     "Test with non-URL",
			norm:     false,
			args:     []any{"test"},
			expected: "test",
		},
		{
			name:     "Test with URL and norm true",
			norm:     true,
			args:     []any{"https://example.com"},
			expected: "https://example.com",
		},
		{
			name:     "Test with non-URL and norm true",
			norm:     true,
			args:     []any{"test"},
			expected: "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := urlProcessor(tt.norm, tt.args...)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
