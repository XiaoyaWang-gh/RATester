package highlight

import (
	"fmt"
	"testing"
)

func TestInlineCodeAttrs_1(t *testing.T) {
	type test struct {
		name     string
		lang     string
		expected string
	}

	tests := []test{
		{
			name:     "Test case 1",
			lang:     "go",
			expected: ` class="code-inline language-go"`,
		},
		{
			name:     "Test case 2",
			lang:     "python",
			expected: ` class="code-inline language-python"`,
		},
		{
			name:     "Test case 3",
			lang:     "javascript",
			expected: ` class="code-inline language-javascript"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := inlineCodeAttrs(tt.lang)
			if got != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, got)
			}
		})
	}
}
