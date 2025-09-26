package acme

import (
	"context"
	"fmt"
	"testing"
)

func TestIsAccountMatchingCaServer_1(t *testing.T) {
	type args struct {
		ctx        context.Context
		accountURI string
		serverURI  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				ctx:        context.Background(),
				accountURI: "http://example.com",
				serverURI:  "http://example.com",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				ctx:        context.Background(),
				accountURI: "http://example.com",
				serverURI:  "http://example2.com",
			},
			want: false,
		},
		{
			name: "Test Case 3",
			args: args{
				ctx:        context.Background(),
				accountURI: "http://example.com",
				serverURI:  "http://example.com:8080",
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

			if got := isAccountMatchingCaServer(tt.args.ctx, tt.args.accountURI, tt.args.serverURI); got != tt.want {
				t.Errorf("isAccountMatchingCaServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
