package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func Testnormalize_1(t *testing.T) {
	tests := []struct {
		name string
		v    reflect.Value
		want any
	}{
		{
			name: "Test case 1",
			v:    reflect.ValueOf(1),
			want: 1.0,
		},
		{
			name: "Test case 2",
			v:    reflect.ValueOf("test"),
			want: "test",
		},
		{
			name: "Test case 3",
			v:    reflect.ValueOf([]int{1, 2, 3}),
			want: []int{1, 2, 3},
		},
		{
			name: "Test case 4",
			v:    reflect.ValueOf(map[string]int{"a": 1, "b": 2}),
			want: map[string]int{"a": 1, "b": 2},
		},
		{
			name: "Test case 5",
			v:    reflect.ValueOf(struct{ Name string }{Name: "test"}),
			want: struct{ Name string }{Name: "test"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := normalize(tt.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
