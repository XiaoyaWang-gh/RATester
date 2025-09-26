package server

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/safe"
	"github.com/traefik/traefik/v3/pkg/server/middleware"
)

func TestNewServer_1(t *testing.T) {
	type args struct {
		routinesPool     *safe.Pool
		entryPoints      TCPEntryPoints
		entryPointsUDP   UDPEntryPoints
		watcher          *ConfigurationWatcher
		observabilityMgr *middleware.ObservabilityMgr
	}
	tests := []struct {
		name string
		args args
		want *Server
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

			if got := NewServer(tt.args.routinesPool, tt.args.entryPoints, tt.args.entryPointsUDP, tt.args.watcher, tt.args.observabilityMgr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
