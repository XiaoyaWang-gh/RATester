package gateway

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestLoadGRPCServers_1(t *testing.T) {
	type args struct {
		namespace  string
		backendRef gatev1.GRPCBackendRef
	}
	tests := []struct {
		name    string
		args    args
		want    *dynamic.ServersLoadBalancer
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

			p := &Provider{}
			got, err := p.loadGRPCServers(tt.args.namespace, tt.args.backendRef)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadGRPCServers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadGRPCServers() = %v, want %v", got, tt.want)
			}
		})
	}
}
