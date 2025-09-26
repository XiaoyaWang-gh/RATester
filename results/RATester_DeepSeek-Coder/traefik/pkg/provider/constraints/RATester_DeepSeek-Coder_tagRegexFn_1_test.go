package constraints

import (
	"fmt"
	"testing"
)

func TestTagRegexFn_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		expr     string
		tags     []string
		expected bool
	}{
		{
			name:     "matching tag",
			expr:     "^tag1$",
			tags:     []string{"tag1", "tag2", "tag3"},
			expected: true,
		},
		{
			name:     "non-matching tag",
			expr:     "^tag4$",
			tags:     []string{"tag1", "tag2", "tag3"},
			expected: false,
		},
		{
			name:     "empty tags",
			expr:     "^tag1$",
			tags:     []string{},
			expected: false,
		},
		{
			name:     "invalid regex",
			expr:     "*",
			tags:     []string{"tag1", "tag2", "tag3"},
			expected: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			fn := tagRegexFn(tt.expr)
			result := fn(tt.tags)
			if result != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
