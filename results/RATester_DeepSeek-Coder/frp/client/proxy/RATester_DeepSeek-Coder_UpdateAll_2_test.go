package proxy

import (
	"context"
	"fmt"
	"net"
	"sync"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/msg"
	"github.com/fatedier/frp/pkg/transport"
)

func TestUpdateAll_2(t *testing.T) {
	type fields struct {
		proxies            map[string]*Wrapper
		msgTransporter     transport.MessageTransporter
		inWorkConnCallback func(*v1.ProxyBaseConfig, net.Conn, *msg.StartWorkConn) bool
		closed             bool
		mu                 sync.RWMutex
		clientCfg          *v1.ClientCommonConfig
		ctx                context.Context
	}
	type args struct {
		proxyCfgs []v1.ProxyConfigurer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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

			pm := &Manager{
				proxies:            tt.fields.proxies,
				msgTransporter:     tt.fields.msgTransporter,
				inWorkConnCallback: tt.fields.inWorkConnCallback,
				closed:             tt.fields.closed,
				mu:                 tt.fields.mu,
				clientCfg:          tt.fields.clientCfg,
				ctx:                tt.fields.ctx,
			}
			pm.UpdateAll(tt.args.proxyCfgs)
		})
	}
}
