package client

import (
	"context"
	"fmt"
	"testing"
)

func TestReload_1(t *testing.T) {
	type args struct {
		ctx        context.Context
		strictMode bool
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		wantErr bool
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

			if err := tt.c.Reload(tt.args.ctx, tt.args.strictMode); (err != nil) != tt.wantErr {
				t.Errorf("Client.Reload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
