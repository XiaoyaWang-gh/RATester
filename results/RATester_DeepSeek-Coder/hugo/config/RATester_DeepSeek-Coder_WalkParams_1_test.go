package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestWalkParams_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		root     maps.Params
		expected []string
	}{
		{
			name: "simple",
			root: maps.Params{
				"a": "1",
				"b": maps.Params{
					"b1": "2",
					"b2": "3",
				},
			},
			expected: []string{"", "a", "b", "b1", "b2"},
		},
		{
			name: "complex",
			root: maps.Params{
				"a": "1",
				"b": maps.Params{
					"b1": "2",
					"b2": maps.Params{
						"b21": "3",
						"b22": "4",
					},
				},
				"c": "5",
			},
			expected: []string{"", "a", "b", "b1", "b2", "b21", "b22", "c"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			provider := &defaultConfigProvider{
				root: test.root,
			}

			var keys []string
			provider.WalkParams(func(params ...maps.KeyParams) bool {
				for _, param := range params {
					keys = append(keys, param.Key)
				}
				return false
			})

			if !reflect.DeepEqual(keys, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, keys)
			}
		})
	}
}
