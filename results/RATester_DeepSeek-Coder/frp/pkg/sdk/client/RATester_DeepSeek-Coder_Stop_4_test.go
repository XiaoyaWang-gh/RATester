package client

import (
	"context"
	"fmt"
	"testing"
)

func TestStop_4(t *testing.T) {
	type args struct {
		ctx context.Context
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

			c := &Client{
				address:  "127.0.0.1:6000",
				authUser: "admin",
				authPwd:  "admin",
			}
			if err := c.Stop(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Client.Stop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
