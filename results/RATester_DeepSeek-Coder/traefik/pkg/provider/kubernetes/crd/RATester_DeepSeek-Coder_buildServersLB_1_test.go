package crd

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
)

func TestBuildServersLB_1(t *testing.T) {
	type args struct {
		namespace string
		svc       traefikv1alpha1.LoadBalancerSpec
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

			c := configBuilder{}
			got, err := c.buildServersLB(tt.args.namespace, tt.args.svc)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildServersLB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildServersLB() = %v, want %v", got, tt.want)
			}
		})
	}
}
