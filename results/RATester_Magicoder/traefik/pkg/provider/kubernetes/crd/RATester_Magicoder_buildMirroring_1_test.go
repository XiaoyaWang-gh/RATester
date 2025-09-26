package crd

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
)

func TestBuildMirroring_1(t *testing.T) {
	type args struct {
		ctx      context.Context
		tService *traefikv1alpha1.TraefikService
		id       string
		conf     map[string]*dynamic.Service
	}
	tests := []struct {
		name    string
		c       configBuilder
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

			if err := tt.c.buildMirroring(tt.args.ctx, tt.args.tService, tt.args.id, tt.args.conf); (err != nil) != tt.wantErr {
				t.Errorf("configBuilder.buildMirroring() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
