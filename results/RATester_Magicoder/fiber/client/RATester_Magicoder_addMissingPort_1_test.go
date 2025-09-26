package client

import (
	"fmt"
	"testing"
)

func TestaddMissingPort_1(t *testing.T) {
	type args struct {
		addr  string
		isTLS bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				addr:  "localhost",
				isTLS: false,
			},
			want: "localhost:80",
		},
		{
			name: "Test Case 2",
			args: args{
				addr:  "localhost",
				isTLS: true,
			},
			want: "localhost:443",
		},
		{
			name: "Test Case 3",
			args: args{
				addr:  "localhost:8080",
				isTLS: false,
			},
			want: "localhost:8080",
		},
		{
			name: "Test Case 4",
			args: args{
				addr:  "localhost:8080",
				isTLS: true,
			},
			want: "localhost:8080",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := addMissingPort(tt.args.addr, tt.args.isTLS); got != tt.want {
				t.Errorf("addMissingPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
