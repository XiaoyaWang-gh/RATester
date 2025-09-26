package main

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/server"
)

func TestSwitchRouter_1(t *testing.T) {
	type args struct {
		routerFactory        *server.RouterFactory
		serverEntryPointsTCP server.TCPEntryPoints
		serverEntryPointsUDP server.UDPEntryPoints
		conf                 dynamic.Configuration
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

			switchRouter(tt.args.routerFactory, tt.args.serverEntryPointsTCP, tt.args.serverEntryPointsUDP)(tt.args.conf)
		})
	}
}
