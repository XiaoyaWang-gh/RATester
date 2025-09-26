package virtual

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/client"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	netpkg "github.com/fatedier/frp/pkg/util/net"
)

func TestUpdateProxyConfigurer_1(t *testing.T) {
	type fields struct {
		l   *netpkg.InternalListener
		svr *client.Service
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

			c := &Client{
				l:   tt.fields.l,
				svr: tt.fields.svr,
			}
			c.UpdateProxyConfigurer(tt.args.proxyCfgs)
		})
	}
}
