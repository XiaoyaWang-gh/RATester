package captcha

import (
	"fmt"
	"testing"
)

func TestRandIntn_1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{n: 10},
			want: 0,
		},
		{
			name: "Test case 2",
			args: args{n: 100},
			want: 0,
		},
		{
			name: "Test case 3",
			args: args{n: 1000},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := randIntn(tt.args.n); got != tt.want {
				t.Errorf("randIntn() = %v, want %v", got, tt.want)
			}
		})
	}
}
