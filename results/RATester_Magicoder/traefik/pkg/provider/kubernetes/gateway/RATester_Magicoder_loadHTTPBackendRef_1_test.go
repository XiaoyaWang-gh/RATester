package gateway

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestLoadHTTPBackendRef_1(t *testing.T) {
	type args struct {
		namespace  string
		backendRef gatev1.HTTPBackendRef
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   *dynamic.Service
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
			got, got1, err := p.loadHTTPBackendRef(tt.args.namespace, tt.args.backendRef)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.loadHTTPBackendRef() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Provider.loadHTTPBackendRef() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Provider.loadHTTPBackendRef() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
