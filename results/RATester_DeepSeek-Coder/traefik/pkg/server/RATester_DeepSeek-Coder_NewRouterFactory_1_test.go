package server

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
	"github.com/traefik/traefik/v3/pkg/server/middleware"
	"github.com/traefik/traefik/v3/pkg/server/service"
	"github.com/traefik/traefik/v3/pkg/tcp"
	"github.com/traefik/traefik/v3/pkg/tls"
)

func TestNewRouterFactory_1(t *testing.T) {
	type args struct {
		staticConfiguration static.Configuration
		managerFactory      *service.ManagerFactory
		tlsManager          *tls.Manager
		observabilityMgr    *middleware.ObservabilityMgr
		pluginBuilder       middleware.PluginsBuilder
		dialerManager       *tcp.DialerManager
	}
	tests := []struct {
		name string
		args args
		want *RouterFactory
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

			if got := NewRouterFactory(tt.args.staticConfiguration, tt.args.managerFactory, tt.args.tlsManager, tt.args.observabilityMgr, tt.args.pluginBuilder, tt.args.dialerManager); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouterFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
