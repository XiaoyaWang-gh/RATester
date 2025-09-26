package crd

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
)

func TestCreateChainMiddleware_1(t *testing.T) {
	type args struct {
		ctx       context.Context
		namespace string
		chain     *traefikv1alpha1.Chain
	}
	tests := []struct {
		name string
		args args
		want *dynamic.Chain
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

			if got := createChainMiddleware(tt.args.ctx, tt.args.namespace, tt.args.chain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createChainMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
