package client

import (
	"context"
	"fmt"
	"net"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestRealConnect_1(t *testing.T) {
	type fields struct {
		ctx context.Context
		cfg *v1.ClientCommonConfig
	}
	tests := []struct {
		name    string
		fields  fields
		want    net.Conn
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

			c := &defaultConnectorImpl{
				ctx: tt.fields.ctx,
				cfg: tt.fields.cfg,
			}
			got, err := c.realConnect()
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultConnectorImpl.realConnect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultConnectorImpl.realConnect() = %v, want %v", got, tt.want)
			}
		})
	}
}
