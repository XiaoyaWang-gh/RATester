package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnewSliceElement_1(t *testing.T) {
	type args struct {
		items any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "Test with slice",
			args: args{
				items: []int{1, 2, 3},
			},
			want: new(int),
		},
		{
			name: "Test with array",
			args: args{
				items: [3]int{1, 2, 3},
			},
			want: new(int),
		},
		{
			name: "Test with nil",
			args: args{
				items: nil,
			},
			want: nil,
		},
		{
			name: "Test with non-slice/array",
			args: args{
				items: 1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := newSliceElement(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSliceElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
