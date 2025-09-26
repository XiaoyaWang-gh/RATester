package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestGetStringMapString_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		key      string
		value    any
		expected map[string]string
	}{
		{
			name:  "returns map[string]string when value is map[string]any",
			key:   "key1",
			value: map[string]any{"k1": "v1", "k2": "v2"},
			expected: map[string]string{
				"k1": "v1",
				"k2": "v2",
			},
		},
		{
			name:     "returns empty map[string]string when value is not map[string]any",
			key:      "key2",
			value:    "not a map",
			expected: map[string]string{},
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

			c := &defaultConfigProvider{
				root: maps.Params{tt.key: tt.value},
			}

			result := c.GetStringMapString(tt.key)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
