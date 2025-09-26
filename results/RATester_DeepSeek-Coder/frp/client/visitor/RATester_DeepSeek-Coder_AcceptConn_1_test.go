package visitor

import (
	"context"
	"fmt"
	"net"
	"sync"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	netpkg "github.com/fatedier/frp/pkg/util/net"
)

func TestAcceptConn_1(t *testing.T) {
	type fields struct {
		clientCfg  *v1.ClientCommonConfig
		helper     Helper
		l          net.Listener
		internalLn *netpkg.InternalListener
		mu         sync.RWMutex
		ctx        context.Context
	}
	type args struct {
		conn net.Conn
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

			v := &BaseVisitor{
				clientCfg:  tt.fields.clientCfg,
				helper:     tt.fields.helper,
				l:          tt.fields.l,
				internalLn: tt.fields.internalLn,
				mu:         tt.fields.mu,
				ctx:        tt.fields.ctx,
			}
			if err := v.AcceptConn(tt.args.conn); (err != nil) != tt.wantErr {
				t.Errorf("BaseVisitor.AcceptConn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
