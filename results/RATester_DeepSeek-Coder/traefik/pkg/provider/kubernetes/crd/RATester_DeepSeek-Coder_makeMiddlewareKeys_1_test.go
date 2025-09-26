package crd

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
)

func TestMakeMiddlewareKeys_1(t *testing.T) {
	type args struct {
		ctx               context.Context
		ingRouteNamespace string
		middlewares       []traefikv1alpha1.MiddlewareRef
	}
	tests := []struct {
		name    string
		args    args
		want    []string
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
			got, err := p.makeMiddlewareKeys(tt.args.ctx, tt.args.ingRouteNamespace, tt.args.middlewares)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.makeMiddlewareKeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Provider.makeMiddlewareKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
