package captcha

import (
	"fmt"
	"testing"
)

func TestRandInt_1(t *testing.T) {
	type args struct {
		from int
		to   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{
				from: 1,
				to:   10,
			},
			want: 1,
		},
		{
			name: "Test case 2",
			args: args{
				from: 100,
				to:   200,
			},
			want: 100,
		},
		{
			name: "Test case 3",
			args: args{
				from: 0,
				to:   0,
			},
			want: 0,
		},
		{
			name: "Test case 4",
			args: args{
				from: -10,
				to:   10,
			},
			want: -10,
		},
		{
			name: "Test case 5",
			args: args{
				from: 10,
				to:   1,
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

			if got := randInt(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("randInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
