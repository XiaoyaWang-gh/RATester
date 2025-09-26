package predicate

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFilterCopy_1(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		p    P[int]
		args args
		want []int
	}{
		{
			name: "Test with positive numbers",
			p:    func(i int) bool { return i > 0 },
			args: args{s: []int{1, 2, -3, 4, -5}},
			want: []int{1, 2, 4},
		},
		{
			name: "Test with negative numbers",
			p:    func(i int) bool { return i < 0 },
			args: args{s: []int{1, 2, -3, 4, -5}},
			want: []int{-3, -5},
		},
		{
			name: "Test with zero",
			p:    func(i int) bool { return i == 0 },
			args: args{s: []int{1, 2, -3, 4, -5}},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.p.FilterCopy(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}
