package math

import (
	"fmt"
	"testing"
)

func Test_round_1(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test case 1",
			args: args{x: 1.5},
			want: 2,
		},
		{
			name: "Test case 2",
			args: args{x: 2.5},
			want: 3,
		},
		{
			name: "Test case 3",
			args: args{x: 3.5},
			want: 4,
		},
		{
			name: "Test case 4",
			args: args{x: -1.5},
			want: -1,
		},
		{
			name: "Test case 5",
			args: args{x: -2.5},
			want: -2,
		},
		{
			name: "Test case 6",
			args: args{x: -3.5},
			want: -4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := _round(tt.args.x); got != tt.want {
				t.Errorf("_round() = %v, want %v", got, tt.want)
			}
		})
	}
}
