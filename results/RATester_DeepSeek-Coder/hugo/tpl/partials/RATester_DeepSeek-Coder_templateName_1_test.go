package partials

import (
	"fmt"
	"testing"
)

func TestTemplateName_1(t *testing.T) {
	tests := []struct {
		name     string
		k        partialCacheKey
		expected string
	}{
		{
			name: "Test with prefix",
			k: partialCacheKey{
				Name: "partials/test",
			},
			expected: "partials/test",
		},
		{
			name: "Test without prefix",
			k: partialCacheKey{
				Name: "test",
			},
			expected: "partials/test",
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
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
