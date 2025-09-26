package nomad

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestAddServerUDP_3(t *testing.T) {
	type args struct {
		i  item
		lb *dynamic.UDPServersLoadBalancer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should return error when load-balancer is missing",
			args: args{
				i: item{
					Address: "127.0.0.1",
					Port:    8080,
				},
				lb: nil,
			},
			wantErr: true,
		},
		{
			name: "should return error when address is missing",
			args: args{
				i: item{
					Port: 8080,
				},
				lb: &dynamic.UDPServersLoadBalancer{},
			},
			wantErr: true,
		},
		{
			name: "should return error when port is missing",
			args: args{
				i: item{
					Address: "127.0.0.1",
				},
				lb: &dynamic.UDPServersLoadBalancer{},
			},
			wantErr: true,
		},
		{
			name: "should add server to load-balancer",
			args: args{
				i: item{
					Address: "127.0.0.1",
					Port:    8080,
				},
				lb: &dynamic.UDPServersLoadBalancer{},
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

			p := &Provider{}
			if err := p.addServerUDP(tt.args.i, tt.args.lb); (err != nil) != tt.wantErr {
				t.Errorf("addServerUDP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
