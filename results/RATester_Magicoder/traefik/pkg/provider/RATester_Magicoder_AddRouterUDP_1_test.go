package provider

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestAddRouterUDP_1(t *testing.T) {
	type args struct {
		configuration *dynamic.UDPConfiguration
		routerName    string
		router        *dynamic.UDPRouter
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "add router",
			args: args{
				configuration: &dynamic.UDPConfiguration{
					Routers: map[string]*dynamic.UDPRouter{},
				},
				routerName: "router1",
				router: &dynamic.UDPRouter{
					EntryPoints: []string{"entry1"},
					Service:     "service1",
				},
			},
			want: true,
		},
		{
			name: "update router",
			args: args{
				configuration: &dynamic.UDPConfiguration{
					Routers: map[string]*dynamic.UDPRouter{
						"router1": {
							EntryPoints: []string{"entry1"},
							Service:     "service1",
						},
					},
				},
				routerName: "router1",
				router: &dynamic.UDPRouter{
					EntryPoints: []string{"entry1"},
					Service:     "service1",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := AddRouterUDP(tt.args.configuration, tt.args.routerName, tt.args.router); got != tt.want {
				t.Errorf("AddRouterUDP() = %v, want %v", got, tt.want)
			}
		})
	}
}
