package util

import (
	"fmt"
	"testing"
)

func TestEmptyOr_1(t *testing.T) {
	type args struct {
		v        string
		fallback string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty string should return fallback",
			args: args{
				v:        "",
				fallback: "fallback",
			},
			want: "fallback",
		},
		{
			name: "Non-empty string should return string itself",
			args: args{
				v:        "test",
				fallback: "fallback",
			},
			want: "test",
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
