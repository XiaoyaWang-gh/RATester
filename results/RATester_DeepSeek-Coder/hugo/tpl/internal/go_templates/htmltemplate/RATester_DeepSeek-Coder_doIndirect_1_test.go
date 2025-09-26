package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDoIndirect_1(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "Test with nil",
			args: args{
				a: nil,
			},
			want: nil,
		},
		{
			name: "Test with non-pointer",
			args: args{
				a: 123,
			},
			want: 123,
		},
		{
			name: "Test with pointer",
			args: args{
				a: &[]int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "Test with double pointer",
			args: args{
				a: &[]*int{&[]int{1, 2, 3}[0]},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := doIndirect(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doIndirect() = %v, want %v", got, tt.want)
			}
		})
	}
}
