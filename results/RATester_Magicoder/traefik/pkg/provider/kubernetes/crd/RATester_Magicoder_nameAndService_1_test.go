package crd

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
)

func TestNameAndService_1(t *testing.T) {
	type args struct {
		ctx             context.Context
		parentNamespace string
		service         traefikv1alpha1.LoadBalancerSpec
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

			c := configBuilder{}
			got, got1, err := c.nameAndService(tt.args.ctx, tt.args.parentNamespace, tt.args.service)
			if (err != nil) != tt.wantErr {
				t.Errorf("nameAndService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("nameAndService() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("nameAndService() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
