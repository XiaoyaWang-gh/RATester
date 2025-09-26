package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestWalkParams_1(t *testing.T) {
	type testCase struct {
		name     string
		root     maps.Params
		expected []string
	}

	testCases := []testCase{
		{
			name: "simple",
			root: maps.Params{
				"a": "1",
				"b": maps.Params{
					"c": "2",
					"d": "3",
				},
			},
			expected: []string{"a", "b", "c", "d"},
		},
		{
			name: "nested",
			root: maps.Params{
				"a": "1",
				"b": maps.Params{
					"c": "2",
					"d": maps.Params{
						"e": "3",
						"f": "4",
					},
				},
			},
			expected: []string{"a", "b", "c", "d", "e", "f"},
		},
		{
			name:     "empty",
			root:     maps.Params{},
			expected: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			provider := &defaultConfigProvider{root: tc.root}
			var keys []string
			provider.WalkParams(func(params ...maps.KeyParams) bool {
				for _, param := range params {
					keys = append(keys, param.Key)
				}
				return false
			})
			if !reflect.DeepEqual(keys, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, keys)
			}
		})
	}
}
