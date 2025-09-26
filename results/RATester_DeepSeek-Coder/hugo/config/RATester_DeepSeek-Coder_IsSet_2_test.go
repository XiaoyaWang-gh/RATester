package config

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestIsSet_2(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		key      string
		expected bool
	}{
		{
			name:     "key exists",
			key:      "existing_key",
			expected: true,
		},
		{
			name:     "key does not exist",
			key:      "non_existing_key",
			expected: false,
		},
	}

	provider := &defaultConfigProvider{
		root: maps.Params{
			"existing_key": "value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := provider.IsSet(tt.key)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
