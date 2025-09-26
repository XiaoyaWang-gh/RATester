package client

import (
	"context"
	"fmt"
	"testing"
)

func TestGetConfig_1(t *testing.T) {
	type fields struct {
		address  string
		authUser string
		authPwd  string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
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
				address:  tt.fields.address,
				authUser: tt.fields.authUser,
				authPwd:  tt.fields.authPwd,
			}
			got, err := c.GetConfig(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
