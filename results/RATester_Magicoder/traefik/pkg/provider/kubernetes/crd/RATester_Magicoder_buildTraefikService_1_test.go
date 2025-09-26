package crd

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
)

func TestBuildTraefikService_1(t *testing.T) {
	type args struct {
		ctx      context.Context
		tService *traefikv1alpha1.TraefikService
		conf     map[string]*dynamic.Service
	}
	tests := []struct {
		name    string
		args    args
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
			if err := c.buildTraefikService(tt.args.ctx, tt.args.tService, tt.args.conf); (err != nil) != tt.wantErr {
				t.Errorf("buildTraefikService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
