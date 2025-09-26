package metadecoders

import (
	"fmt"
	"testing"
)

func TestisLowerIndexThan_1(t *testing.T) {
	type args struct {
		first  int
		others []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				first:  1,
				others: []int{2, 3, 4},
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				first:  5,
				others: []int{2, 3, 4},
			},
			want: false,
		},
		{
			name: "Test Case 3",
			args: args{
				first:  -1,
				others: []int{2, 3, 4},
			},
			want: false,
		},
		{
			name: "Test Case 4",
			args: args{
				first:  1,
				others: []int{-1, 3, 4},
			},
			want: false,
		},
		{
			name: "Test Case 5",
			args: args{
				first:  1,
				others: []int{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isLowerIndexThan(tt.args.first, tt.args.others...); got != tt.want {
				t.Errorf("isLowerIndexThan() = %v, want %v", got, tt.want)
			}
		})
	}
}
