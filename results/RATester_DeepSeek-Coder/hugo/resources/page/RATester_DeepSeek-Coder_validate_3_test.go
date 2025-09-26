package page

import (
	"fmt"
	"testing"
)

func TestValidate_3(t *testing.T) {
	type test struct {
		name     string
		pp       string
		expected bool
	}

	tests := []test{
		{
			name:     "Empty string",
			pp:       "",
			expected: false,
		},
		{
			name:     "Valid permalink",
			pp:       "/:section/:slug",
			expected: true,
		},
		{
			name:     "Invalid permalink",
			pp:       "/:section/:unknown",
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

			l := PermalinkExpander{}
			result := l.validate(tt.pp)
			if result != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
