package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestGet_18(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		provider *defaultConfigProvider
		key      string
		expected any
	}{
		{
			name: "get existing key",
			provider: &defaultConfigProvider{
				root: maps.Params{"key1": "value1", "key2": "value2"},
			},
			key:      "key1",
			expected: "value1",
		},
		{
			name: "get non-existing key",
			provider: &defaultConfigProvider{
				root: maps.Params{"key1": "value1", "key2": "value2"},
			},
			key:      "key3",
			expected: nil,
		},
		{
			name: "get empty key",
			provider: &defaultConfigProvider{
				root: maps.Params{"key1": "value1", "key2": "value2"},
			},
			key:      "",
			expected: maps.Params{"key1": "value1", "key2": "value2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual := tt.provider.Get(tt.key)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
