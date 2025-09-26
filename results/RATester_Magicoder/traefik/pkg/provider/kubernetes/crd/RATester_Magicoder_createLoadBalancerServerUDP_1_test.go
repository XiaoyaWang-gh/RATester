package crd

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
)

func TestCreateLoadBalancerServerUDP_1(t *testing.T) {
	type args struct {
		client          Client
		parentNamespace string
		service         traefikv1alpha1.ServiceUDP
	}
	tests := []struct {
		name    string
		args    args
		want    *dynamic.UDPService
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
			got, err := p.createLoadBalancerServerUDP(tt.args.client, tt.args.parentNamespace, tt.args.service)
			if (err != nil) != tt.wantErr {
				t.Errorf("createLoadBalancerServerUDP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createLoadBalancerServerUDP() = %v, want %v", got, tt.want)
			}
		})
	}
}
