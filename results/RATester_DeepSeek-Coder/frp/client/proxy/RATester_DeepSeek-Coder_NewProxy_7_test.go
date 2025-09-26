package proxy

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/transport"
)

func TestNewProxy_7(t *testing.T) {
	type args struct {
		ctx            context.Context
		pxyConf        v1.ProxyConfigurer
		clientCfg      *v1.ClientCommonConfig
		msgTransporter transport.MessageTransporter
	}
	tests := []struct {
		name    string
		args    args
		wantPxy Proxy
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

			gotPxy := NewProxy(tt.args.ctx, tt.args.pxyConf, tt.args.clientCfg, tt.args.msgTransporter)
			if !reflect.DeepEqual(gotPxy, tt.wantPxy) {
				t.Errorf("NewProxy() = %v, want %v", gotPxy, tt.wantPxy)
			}
		})
	}
}
