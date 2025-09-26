package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestGetStringSlice_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		provider *defaultConfigProvider
		key      string
		expected []string
	}{
		{
			name: "Test with existing key",
			provider: &defaultConfigProvider{
				root: maps.Params{"key1": []string{"value1", "value2"}},
			},
			key:      "key1",
			expected: []string{"value1", "value2"},
		},
		{
			name: "Test with non-existing key",
			provider: &defaultConfigProvider{
				root: maps.Params{"key2": "value3"},
			},
			key:      "key1",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tt.provider.GetStringSlice(tt.key)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
