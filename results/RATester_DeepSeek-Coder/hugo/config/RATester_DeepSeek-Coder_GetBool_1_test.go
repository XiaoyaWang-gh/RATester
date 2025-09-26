package config

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestGetBool_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		key      string
		value    any
		expected bool
	}{
		{
			name:     "GetBool returns true for key with true value",
			key:      "key1",
			value:    true,
			expected: true,
		},
		{
			name:     "GetBool returns false for key with false value",
			key:      "key2",
			value:    false,
			expected: false,
		},
		{
			name:     "GetBool returns false for key with non-bool value",
			key:      "key3",
			value:    "non-bool",
			expected: false,
		},
		{
			name:     "GetBool returns false for key with nil value",
			key:      "key4",
			value:    nil,
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

			provider := &defaultConfigProvider{
				root: maps.Params{tt.key: tt.value},
			}

			result := provider.GetBool(tt.key)

			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
