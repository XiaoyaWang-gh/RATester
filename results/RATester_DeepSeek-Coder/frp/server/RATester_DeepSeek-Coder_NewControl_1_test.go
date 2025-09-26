package server

import (
	"context"
	"fmt"
	"net"
	"reflect"
	"testing"

	"github.com/fatedier/frp/pkg/auth"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/msg"
	plugin "github.com/fatedier/frp/pkg/plugin/server"
	"github.com/fatedier/frp/server/controller"
	"github.com/fatedier/frp/server/proxy"
)

func TestNewControl_1(t *testing.T) {
	type args struct {
		ctx              context.Context
		rc               *controller.ResourceController
		pxyManager       *proxy.Manager
		pluginManager    *plugin.Manager
		authVerifier     auth.Verifier
		ctlConn          net.Conn
		ctlConnEncrypted bool
		loginMsg         *msg.Login
		serverCfg        *v1.ServerConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *Control
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

			got, err := NewControl(tt.args.ctx, tt.args.rc, tt.args.pxyManager, tt.args.pluginManager, tt.args.authVerifier, tt.args.ctlConn, tt.args.ctlConnEncrypted, tt.args.loginMsg, tt.args.serverCfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewControl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewControl() = %v, want %v", got, tt.want)
			}
		})
	}
}
