package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestGetParams_1(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		key      string
		expected maps.Params
	}

	tests := []testCase{
		{
			name:     "existing key",
			key:      "existingKey",
			expected: maps.Params{"key1": "value1", "key2": "value2"},
		},
		{
			name:     "non-existing key",
			key:      "nonExistingKey",
			expected: nil,
		},
	}

	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			provider := &defaultConfigProvider{
				root: maps.Params{"existingKey": tc.expected},
			}

			result := provider.GetParams(tc.key)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
