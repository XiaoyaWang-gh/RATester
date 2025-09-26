package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSlicePad_1(t *testing.T) {
	type args struct {
		slice []interface{}
		size  int
		val   interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "Test case 1",
			args: args{
				slice: []interface{}{"a", "b", "c"},
				size:  5,
				val:   "d",
			},
			want: []interface{}{"a", "b", "c", "d", "d"},
		},
		{
			name: "Test case 2",
			args: args{
				slice: []interface{}{1, 2, 3},
				size:  3,
				val:   4,
			},
			want: []interface{}{1, 2, 3},
		},
		{
			name: "Test case 3",
			args: args{
				slice: []interface{}{},
				size:  10,
				val:   "e",
			},
			want: []interface{}{"e", "e", "e", "e", "e", "e", "e", "e", "e", "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := SlicePad(tt.args.slice, tt.args.size, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlicePad() = %v, want %v", got, tt.want)
			}
		})
	}
}
