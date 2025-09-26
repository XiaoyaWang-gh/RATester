package util

import (
	"fmt"
	"testing"
)

func TestEmptyOr_1(t *testing.T) {
	type args struct {
		v        int
		fallback int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{
				v:        0,
				fallback: 1,
			},
			want: 1,
		},
		{
			name: "Test case 2",
			args: args{
				v:        2,
				fallback: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := EmptyOr(tt.args.v, tt.args.fallback); got != tt.want {
				t.Errorf("EmptyOr() = %v, want %v", got, tt.want)
			}
		})
	}
}
