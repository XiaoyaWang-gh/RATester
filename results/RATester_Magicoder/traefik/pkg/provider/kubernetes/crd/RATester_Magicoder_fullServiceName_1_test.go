package crd

import (
	"context"
	"fmt"
	"testing"

	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func TestFullServiceName_1(t *testing.T) {
	type args struct {
		ctx       context.Context
		namespace string
		service   traefikv1alpha1.LoadBalancerSpec
		port      intstr.IntOrString
	}
	tests := []struct {
		name string
		args args
		want string
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

			if got := fullServiceName(tt.args.ctx, tt.args.namespace, tt.args.service, tt.args.port); got != tt.want {
				t.Errorf("fullServiceName() = %v, want %v", got, tt.want)
			}
		})
	}
}
