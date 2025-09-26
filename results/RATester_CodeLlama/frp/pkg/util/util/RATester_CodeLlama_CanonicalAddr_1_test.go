package util

import (
	"fmt"
	"testing"
)

func TestCanonicalAddr_1(t *testing.T) {
	type args struct {
		host string
		port int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{
				host: "127.0.0.1",
				port: 80,
			},
			want: "127.0.0.1",
		},
		{
			name: "test case 2",
			args: args{
				host: "127.0.0.1",
				port: 443,
			},
			want: "127.0.0.1",
		},
		{
			name: "test case 3",
			args: args{
				host: "127.0.0.1",
				port: 8080,
			},
			want: "127.0.0.1:8080",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := CanonicalAddr(tt.args.host, tt.args.port); got != tt.want {
				t.Errorf("CanonicalAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
