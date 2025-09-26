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
		// TODO: Add test cases.
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
