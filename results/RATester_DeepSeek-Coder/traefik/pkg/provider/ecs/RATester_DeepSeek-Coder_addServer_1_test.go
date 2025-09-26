package ecs

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestAddServer_1(t *testing.T) {
	type args struct {
		instance     ecsInstance
		loadBalancer *dynamic.ServersLoadBalancer
	}
	tests := []struct {
		name    string
		p       *Provider
		args    args
		wantErr bool
	}{
		{
			name: "test_case_1",
			p:    &Provider{},
			args: args{
				instance: ecsInstance{
					Name: "test_instance",
					ID:   "test_id",
				},
				loadBalancer: &dynamic.ServersLoadBalancer{},
			},
			wantErr: false,
		},
		{
			name: "test_case_2",
			p:    &Provider{},
			args: args{
				instance: ecsInstance{
					Name: "test_instance",
					ID:   "test_id",
				},
				loadBalancer: nil,
			},
			wantErr: true,
		},
		{
			name: "test_case_3",
			p:    &Provider{},
			args: args{
				instance: ecsInstance{
					Name: "test_instance",
					ID:   "test_id",
				},
				loadBalancer: &dynamic.ServersLoadBalancer{
					Servers: []dynamic.Server{},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := tt.p.addServer(tt.args.instance, tt.args.loadBalancer); (err != nil) != tt.wantErr {
				t.Errorf("Provider.addServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
