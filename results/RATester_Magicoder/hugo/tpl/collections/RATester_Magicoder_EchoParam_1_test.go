package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEchoParam_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name string
		c    any
		k    any
		want any
	}{
		{
			name: "Test with array",
			c:    []int{1, 2, 3, 4, 5},
			k:    2,
			want: 3,
		},
		{
			name: "Test with slice",
			c:    []int{1, 2, 3, 4, 5},
			k:    2,
			want: 3,
		},
		{
			name: "Test with map",
			c:    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			k:    "c",
			want: 3,
		},
		{
			name: "Test with nil",
			c:    nil,
			k:    2,
			want: "",
		},
		{
			name: "Test with invalid type",
			c:    []int{1, 2, 3, 4, 5},
			k:    "2",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ns.EchoParam(tt.c, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EchoParam() = %v, want %v", got, tt.want)
			}
		})
	}
}
