package client

import (
	"context"
	"fmt"
	"net"
	"reflect"
	"sync"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	fmux "github.com/hashicorp/yamux"
	quic "github.com/quic-go/quic-go"
)

func TestConnect_1(t *testing.T) {
	t.Parallel()

	type fields struct {
		ctx        context.Context
		cfg        *v1.ClientCommonConfig
		muxSession *fmux.Session
		quicConn   quic.Connection
		closeOnce  sync.Once
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
				ctx:        tt.fields.ctx,
				cfg:        tt.fields.cfg,
				muxSession: tt.fields.muxSession,
				quicConn:   tt.fields.quicConn,
				closeOnce:  tt.fields.closeOnce,
			}
			got, err := c.Connect()
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultConnectorImpl.Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultConnectorImpl.Connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
