package docker

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestAddServerTCP_2(t *testing.T) {
	type args struct {
		ctx          context.Context
		container    dockerData
		loadBalancer *dynamic.TCPServersLoadBalancer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
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

			p := &DynConfBuilder{}
			if err := p.addServerTCP(tt.args.ctx, tt.args.container, tt.args.loadBalancer); (err != nil) != tt.wantErr {
				t.Errorf("DynConfBuilder.addServerTCP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
