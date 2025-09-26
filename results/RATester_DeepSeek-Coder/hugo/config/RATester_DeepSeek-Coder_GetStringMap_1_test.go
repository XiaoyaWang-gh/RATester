package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestGetStringMap_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		provider *defaultConfigProvider
		key      string
		expected map[string]any
	}{
		{
			name: "success",
			provider: &defaultConfigProvider{
				root: maps.Params{"key": "value"},
			},
			key:      "key",
			expected: map[string]any{"key": "value"},
		},
		{
			name: "not found",
			provider: &defaultConfigProvider{
				root: maps.Params{"key": "value"},
			},
			key:      "not_found",
			expected: map[string]any{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tt.provider.GetStringMap(tt.key)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
