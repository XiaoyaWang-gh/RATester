package visitor

import (
	"fmt"
	"net"
	"sync"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	quic "github.com/quic-go/quic-go"
)

func TestInit_9(t *testing.T) {
	type fields struct {
		session    quic.Connection
		listenConn *net.UDPConn
		mu         sync.RWMutex
		clientCfg  *v1.ClientCommonConfig
	}
	type args struct {
		listenConn *net.UDPConn
		raddr      *net.UDPAddr
	}
	tests := []struct {
		name    string
		fields  fields
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

			qs := &QUICTunnelSession{
				session:    tt.fields.session,
				listenConn: tt.fields.listenConn,
				mu:         tt.fields.mu,
				clientCfg:  tt.fields.clientCfg,
			}
			if err := qs.Init(tt.args.listenConn, tt.args.raddr); (err != nil) != tt.wantErr {
				t.Errorf("QUICTunnelSession.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
