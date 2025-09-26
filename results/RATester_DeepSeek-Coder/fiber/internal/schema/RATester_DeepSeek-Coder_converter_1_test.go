package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConverter_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		cache    *cache
		input    reflect.Type
		expected Converter
	}{
		{
			name: "Test with existing type",
			cache: &cache{
				regconv: map[reflect.Type]Converter{
					reflect.TypeOf(1): func(s string) reflect.Value {
						return reflect.ValueOf(s)
					},
				},
			},
			input: reflect.TypeOf(1),
			expected: func(s string) reflect.Value {
				return reflect.ValueOf(s)
			},
		},
		{
			name: "Test with non-existing type",
			cache: &cache{
				regconv: map[reflect.Type]Converter{
					reflect.TypeOf(1): func(s string) reflect.Value {
						return reflect.ValueOf(s)
					},
				},
			},
			input:    reflect.TypeOf("test"),
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.cache.converter(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
