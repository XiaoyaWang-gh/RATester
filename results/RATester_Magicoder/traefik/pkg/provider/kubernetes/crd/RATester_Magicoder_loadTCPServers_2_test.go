package crd

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
)

func TestLoadTCPServers_2(t *testing.T) {
	type args struct {
		client     Client
		namespace  string
		serviceTCP traefikv1alpha1.ServiceTCP
	}
	tests := []struct {
		name    string
		args    args
		want    []dynamic.TCPServer
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
			got, err := p.loadTCPServers(tt.args.client, tt.args.namespace, tt.args.serviceTCP)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadTCPServers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadTCPServers() = %v, want %v", got, tt.want)
			}
		})
	}
}
