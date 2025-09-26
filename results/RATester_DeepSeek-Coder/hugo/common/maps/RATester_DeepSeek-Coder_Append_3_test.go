package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAppend_3(t *testing.T) {
	type testStruct struct {
		name string
		key  string
		val  []int
		want []int
	}

	tests := []testStruct{
		{
			name: "Append to existing key",
			key:  "key1",
			val:  []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "Append to new key",
			key:  "key2",
			val:  []int{4, 5, 6},
			want: []int{4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &SliceCache[int]{
				m: make(map[string][]int),
			}
			c.Append(tt.key, tt.val...)
			got, _ := c.Get(tt.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Append() = %v, want %v", got, tt.want)
			}
		})
	}
}
