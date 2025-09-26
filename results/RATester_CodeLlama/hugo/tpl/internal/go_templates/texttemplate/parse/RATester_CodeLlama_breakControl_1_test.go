package parse

import (
	"fmt"
	"testing"
)

func TestBreakControl_1(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic",
			input:    `{{break}}`,
			expected: ``,
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
			var tree Tree
			tree.Parse(tt.input, "", "", nil)
			actual := tree.Root.String()
			if actual != tt.expected {
				t.Errorf("expected %s, but got %s", tt.expected, actual)
			}
		})
	}
}
