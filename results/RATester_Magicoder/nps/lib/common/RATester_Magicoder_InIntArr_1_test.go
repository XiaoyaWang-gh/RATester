package common

import (
	"fmt"
	"testing"
)

func TestInIntArr_1(t *testing.T) {
	type args struct {
		arr []int
		val int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				val: 3,
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				val: 6,
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				arr: []int{},
				val: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := InIntArr(tt.args.arr, tt.args.val); got != tt.want {
				t.Errorf("InIntArr() = %v, want %v", got, tt.want)
			}
		})
	}
}
