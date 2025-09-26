package consulcatalog

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestAddServer_5(t *testing.T) {
	type args struct {
		item         itemData
		loadBalancer *dynamic.ServersLoadBalancer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test_case_1",
			args: args{
				item: itemData{
					Address: "127.0.0.1",
					Port:    "80",
				},
				loadBalancer: &dynamic.ServersLoadBalancer{},
			},
			wantErr: false,
		},
		{
			name: "test_case_2",
			args: args{
				item: itemData{
					Address: "",
					Port:    "80",
				},
				loadBalancer: &dynamic.ServersLoadBalancer{},
			},
			wantErr: true,
		},
		{
			name: "test_case_3",
			args: args{
				item: itemData{
					Address: "127.0.0.1",
					Port:    "",
				},
				loadBalancer: &dynamic.ServersLoadBalancer{},
			},
			wantErr: true,
		},
		{
			name: "test_case_4",
			args: args{
				item: itemData{
					Address: "",
					Port:    "",
				},
				loadBalancer: &dynamic.ServersLoadBalancer{},
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
			if err := p.addServer(tt.args.item, tt.args.loadBalancer); (err != nil) != tt.wantErr {
				t.Errorf("addServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
