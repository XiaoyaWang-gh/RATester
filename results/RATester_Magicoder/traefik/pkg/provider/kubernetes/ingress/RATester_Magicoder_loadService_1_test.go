package ingress

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	netv1 "k8s.io/api/networking/v1"
)

func TestLoadService_1(t *testing.T) {
	type args struct {
		client    Client
		namespace string
		backend   netv1.IngressBackend
	}
	tests := []struct {
		name    string
		args    args
		want    *dynamic.Service
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
			got, err := p.loadService(tt.args.client, tt.args.namespace, tt.args.backend)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadService() = %v, want %v", got, tt.want)
			}
		})
	}
}
