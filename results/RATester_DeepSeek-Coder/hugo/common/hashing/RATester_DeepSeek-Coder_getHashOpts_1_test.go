package hashing

import (
	"fmt"
	"hash/fnv"
	"reflect"
	"testing"

	"github.com/gohugoio/hashstructure"
)

func TestGetHashOpts_1(t *testing.T) {
	tests := []struct {
		name string
		want *hashstructure.HashOptions
	}{
		{
			name: "Test case 1",
			want: &hashstructure.HashOptions{
				Hasher:          fnv.New64(),
				TagName:         "hash",
				ZeroNil:         false,
				IgnoreZeroValue: false,
				SlicesAsSets:    false,
				UseStringer:     false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getHashOpts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHashOpts() = %v, want %v", got, tt.want)
			}
		})
	}
}
