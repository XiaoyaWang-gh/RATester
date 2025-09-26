package hashing

import (
	"fmt"
	"hash/fnv"
	"reflect"
	"testing"

	"github.com/gohugoio/hashstructure"
)

func TestPutHashOpts_1(t *testing.T) {
	testCases := []struct {
		name     string
		opts     *hashstructure.HashOptions
		expected *hashstructure.HashOptions
	}{
		{
			name: "putHashOpts with nil opts",
			opts: nil,
			expected: &hashstructure.HashOptions{
				Hasher:          nil,
				TagName:         "",
				ZeroNil:         false,
				IgnoreZeroValue: false,
				SlicesAsSets:    false,
				UseStringer:     false,
			},
		},
		{
			name: "putHashOpts with non-nil opts",
			opts: &hashstructure.HashOptions{
				Hasher:          fnv.New64(),
				TagName:         "hash",
				ZeroNil:         true,
				IgnoreZeroValue: true,
				SlicesAsSets:    true,
				UseStringer:     true,
			},
			expected: &hashstructure.HashOptions{
				Hasher:          fnv.New64(),
				TagName:         "hash",
				ZeroNil:         true,
				IgnoreZeroValue: true,
				SlicesAsSets:    true,
				UseStringer:     true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			putHashOpts(tc.opts)
			if !reflect.DeepEqual(tc.opts, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, tc.opts)
			}
		})
	}
}
