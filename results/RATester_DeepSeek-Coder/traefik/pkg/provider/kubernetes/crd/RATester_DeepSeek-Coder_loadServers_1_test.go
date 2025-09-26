package crd

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
)

func TestLoadServers_1(t *testing.T) {
	type args struct {
		parentNamespace string
		svc             traefikv1alpha1.LoadBalancerSpec
	}
	tests := []struct {
		name    string
		args    args
		want    []dynamic.Server
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

			c := configBuilder{}
			got, err := c.loadServers(tt.args.parentNamespace, tt.args.svc)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadServers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadServers() = %v, want %v", got, tt.want)
			}
		})
	}
}
