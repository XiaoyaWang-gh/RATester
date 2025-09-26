package ecs

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestAddServerTCP_1(t *testing.T) {
	type args struct {
		instance     ecsInstance
		loadBalancer *dynamic.TCPServersLoadBalancer
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
				loadBalancer: &dynamic.TCPServersLoadBalancer{
					Servers: []dynamic.TCPServer{
						{
							Address: "127.0.0.1:8080",
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
				loadBalancer: &dynamic.TCPServersLoadBalancer{
					Servers: []dynamic.TCPServer{},
				},
			},
			wantErr: true,
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
			if err := p.addServerTCP(tt.args.instance, tt.args.loadBalancer); (err != nil) != tt.wantErr {
				t.Errorf("addServerTCP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
