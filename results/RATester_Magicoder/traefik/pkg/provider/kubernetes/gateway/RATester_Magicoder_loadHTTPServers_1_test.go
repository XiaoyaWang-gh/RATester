package gateway

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestLoadHTTPServers_1(t *testing.T) {
	type args struct {
		namespace  string
		backendRef gatev1.HTTPBackendRef
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
			got, err := p.loadHTTPServers(tt.args.namespace, tt.args.backendRef)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadHTTPServers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadHTTPServers() = %v, want %v", got, tt.want)
			}
		})
	}
}
