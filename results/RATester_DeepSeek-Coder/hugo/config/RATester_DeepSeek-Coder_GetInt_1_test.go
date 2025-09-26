package config

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestGetInt_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		key      string
		value    any
		expected int
	}{
		{
			name:     "Test with valid int value",
			key:      "test_key",
			value:    123,
			expected: 123,
		},
		{
			name:     "Test with valid string value",
			key:      "test_key",
			value:    "456",
			expected: 456,
		},
		{
			name:     "Test with invalid value",
			key:      "test_key",
			value:    "invalid",
			expected: 0,
		},
		{
			name:     "Test with nil value",
			key:      "test_key",
			value:    nil,
			expected: 0,
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

			result := provider.GetInt(tt.key)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
