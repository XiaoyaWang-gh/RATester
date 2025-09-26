package crd

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
)

func TestLoadUDPServers_1(t *testing.T) {
	type args struct {
		client    Client
		namespace string
		svc       traefikv1alpha1.ServiceUDP
	}
	tests := []struct {
		name    string
		args    args
		want    []dynamic.UDPServer
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
			got, err := p.loadUDPServers(tt.args.client, tt.args.namespace, tt.args.svc)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadUDPServers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadUDPServers() = %v, want %v", got, tt.want)
			}
		})
	}
}
