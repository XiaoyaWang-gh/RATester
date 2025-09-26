package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestGetParams_1(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected maps.Params
	}

	tests := []testCase{
		{
			name:     "Test case 1",
			key:      "key1",
			expected: maps.Params{"key1": "value1"},
		},
		{
			name:     "Test case 2",
			key:      "key2",
			expected: maps.Params{"key2": "value2"},
		},
		// Add more test cases as needed
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			provider := &defaultConfigProvider{
				root: tc.expected,
			}

			result := provider.GetParams(tc.key)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
