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
		args    args
		wantErr bool
	}{
		{
			name: "test_1",
			args: args{
				instance: ecsInstance{
					Name: "test_instance",
					ID:   "test_id",
				},
				loadBalancer: &dynamic.ServersLoadBalancer{
					Servers: []dynamic.Server{
						{
							URL: "http://test_url",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test_2",
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
		{
			name: "test_3",
			args: args{
				instance: ecsInstance{
					Name: "test_instance",
					ID:   "test_id",
				},
				loadBalancer: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &Provider{}
			if err := p.addServer(tt.args.instance, tt.args.loadBalancer); (err != nil) != tt.wantErr {
				t.Errorf("Provider.addServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
