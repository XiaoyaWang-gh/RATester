package client

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/fatedier/frp/client"
)

func TestGetProxyStatus_2(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *client.ProxyStatusResp
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
				address:  "localhost:6000",
				authUser: "admin",
				authPwd:  "admin",
			}
			got, err := c.GetProxyStatus(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetProxyStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetProxyStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
