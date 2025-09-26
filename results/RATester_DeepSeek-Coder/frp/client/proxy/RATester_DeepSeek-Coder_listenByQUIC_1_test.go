package proxy

import (
	"fmt"
	"net"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/msg"
)

func TestListenByQUIC_1(t *testing.T) {
	type args struct {
		listenConn       *net.UDPConn
		raddr            *net.UDPAddr
		startWorkConnMsg *msg.StartWorkConn
	}
	tests := []struct {
		name string
		args args
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

			pxy := &XTCPProxy{
				BaseProxy: &BaseProxy{
					baseCfg: &v1.ProxyBaseConfig{},
				},
			}
			pxy.listenByQUIC(tt.args.listenConn, tt.args.raddr, tt.args.startWorkConnMsg)
		})
	}
}
