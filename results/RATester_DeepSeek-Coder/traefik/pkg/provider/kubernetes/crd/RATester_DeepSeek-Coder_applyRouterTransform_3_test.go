package crd

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
)

func TestApplyRouterTransform_3(t *testing.T) {
	type args struct {
		ctx     context.Context
		rt      *dynamic.Router
		ingress *traefikv1alpha1.IngressRoute
	}
	tests := []struct {
		name string
		args args
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
			p.applyRouterTransform(tt.args.ctx, tt.args.rt, tt.args.ingress)
		})
	}
}
