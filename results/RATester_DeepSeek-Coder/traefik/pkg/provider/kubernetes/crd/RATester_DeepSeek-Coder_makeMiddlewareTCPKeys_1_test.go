package crd

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
)

func TestMakeMiddlewareTCPKeys_1(t *testing.T) {
	type args struct {
		ctx                  context.Context
		ingRouteTCPNamespace string
		middlewares          []traefikv1alpha1.ObjectReference
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
			got, err := p.makeMiddlewareTCPKeys(tt.args.ctx, tt.args.ingRouteTCPNamespace, tt.args.middlewares)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.makeMiddlewareTCPKeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Provider.makeMiddlewareTCPKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
