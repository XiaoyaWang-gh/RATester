package http

import (
	"fmt"
	"testing"
)

func TestHasPort_1(t *testing.T) {
	type args struct {
		host string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				host: "127.0.0.1:8080",
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				host: "127.0.0.1",
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				host: "[::1]:8080",
			},
			want: true,
		},
		{
			name: "case 4",
			args: args{
				host: "[::1]",
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

			if got := hasPort(tt.args.host); got != tt.want {
				t.Errorf("hasPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
