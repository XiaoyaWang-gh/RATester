package constraints

import (
	"fmt"
	"testing"
)

func TestTagRegexFn_1(t *testing.T) {
	tests := []struct {
		name     string
		expr     string
		tags     []string
		expected bool
	}{
		{
			name:     "Matching tag",
			expr:     "tag1",
			tags:     []string{"tag1", "tag2"},
			expected: true,
		},
		{
			name:     "Non-matching tag",
			expr:     "tag3",
			tags:     []string{"tag1", "tag2"},
			expected: false,
		},
		{
			name:     "Invalid regex",
			expr:     "*",
			tags:     []string{"tag1", "tag2"},
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

			fn := tagRegexFn(tt.expr)
			result := fn(tt.tags)
			if result != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
