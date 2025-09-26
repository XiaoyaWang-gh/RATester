package provider

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestAddTransport_1(t *testing.T) {
	type args struct {
		configuration *dynamic.HTTPConfiguration
		transportName string
		transport     *dynamic.ServersTransport
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "add transport",
			args: args{
				configuration: &dynamic.HTTPConfiguration{
					ServersTransports: make(map[string]*dynamic.ServersTransport),
				},
				transportName: "test",
				transport: &dynamic.ServersTransport{
					ServerName: "test",
				},
			},
			want: true,
		},
		{
			name: "update transport",
			args: args{
				configuration: &dynamic.HTTPConfiguration{
					ServersTransports: map[string]*dynamic.ServersTransport{
						"test": {
							ServerName: "old",
						},
					},
				},
				transportName: "test",
				transport: &dynamic.ServersTransport{
					ServerName: "new",
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

			if got := AddTransport(tt.args.configuration, tt.args.transportName, tt.args.transport); got != tt.want {
				t.Errorf("AddTransport() = %v, want %v", got, tt.want)
			}
		})
	}
}
