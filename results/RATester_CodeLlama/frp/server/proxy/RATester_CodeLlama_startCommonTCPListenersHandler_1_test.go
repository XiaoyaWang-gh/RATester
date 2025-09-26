package proxy

import (
	"context"
	"fmt"
	"net"
	"sync"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/msg"
	plugin "github.com/fatedier/frp/pkg/plugin/server"
	"github.com/fatedier/frp/pkg/util/xlog"
	"github.com/fatedier/frp/server/controller"
	"golang.org/x/time/rate"
)

func TestStartCommonTCPListenersHandler_1(t *testing.T) {
	type fields struct {
		name          string
		rc            *controller.ResourceController
		listeners     []net.Listener
		usedPortsNum  int
		poolCount     int
		getWorkConnFn GetWorkConnFn
		serverCfg     *v1.ServerConfig
		limiter       *rate.Limiter
		userInfo      plugin.UserInfo
		loginMsg      *msg.Login
		configurer    v1.ProxyConfigurer
		mu            sync.RWMutex
		xl            *xlog.Logger
		ctx           context.Context
	}
	tests := []struct {
		name   string
		fields fields
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

			pxy := &BaseProxy{
				name:          tt.fields.name,
				rc:            tt.fields.rc,
				listeners:     tt.fields.listeners,
				usedPortsNum:  tt.fields.usedPortsNum,
				poolCount:     tt.fields.poolCount,
				getWorkConnFn: tt.fields.getWorkConnFn,
				serverCfg:     tt.fields.serverCfg,
				limiter:       tt.fields.limiter,
				userInfo:      tt.fields.userInfo,
				loginMsg:      tt.fields.loginMsg,
				configurer:    tt.fields.configurer,
				mu:            tt.fields.mu,
				xl:            tt.fields.xl,
				ctx:           tt.fields.ctx,
			}
			pxy.startCommonTCPListenersHandler()
		})
	}
}
