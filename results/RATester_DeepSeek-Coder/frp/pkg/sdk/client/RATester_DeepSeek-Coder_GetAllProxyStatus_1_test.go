package client

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/fatedier/frp/client"
)

func TestGetAllProxyStatus_1(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    client.StatusResp
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

			got, err := tt.c.GetAllProxyStatus(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetAllProxyStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetAllProxyStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
