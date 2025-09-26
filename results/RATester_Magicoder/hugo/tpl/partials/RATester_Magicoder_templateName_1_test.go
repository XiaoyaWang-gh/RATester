package partials

import (
	"fmt"
	"testing"
)

func TesttemplateName_1(t *testing.T) {
	tests := []struct {
		name     string
		k        partialCacheKey
		expected string
	}{
		{
			name: "Test case 1",
			k: partialCacheKey{
				Name: "partials/test.tmpl",
			},
			expected: "partials/test.tmpl",
		},
		{
			name: "Test case 2",
			k: partialCacheKey{
				Name: "test.tmpl",
			},
			expected: "partials/test.tmpl",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tt.k.templateName()
			if result != tt.expected {
				t.Errorf("Expected %s, but got %s", tt.expected, result)
			}
		})
	}
}
