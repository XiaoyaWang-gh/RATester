package utils

import (
	"fmt"
	"testing"
)

func TestGet_20(t *testing.T) {
	type args struct {
		i     int
		args  []int
		input ArgInt
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{
				i:     0,
				args:  []int{10},
				input: ArgInt{1, 2, 3},
			},
			want: 1,
		},
		{
			name: "Test case 2",
			args: args{
				i:     2,
				args:  []int{10},
				input: ArgInt{1, 2, 3},
			},
			want: 10,
		},
		{
			name: "Test case 3",
			args: args{
				i:     3,
				args:  []int{10},
				input: ArgInt{1, 2, 3},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.args.input.Get(tt.args.i, tt.args.args...); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
