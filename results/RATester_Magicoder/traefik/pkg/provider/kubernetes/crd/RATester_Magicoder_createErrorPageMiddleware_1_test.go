package crd

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
)

func TestCreateErrorPageMiddleware_1(t *testing.T) {
	type args struct {
		client    Client
		namespace string
		errorPage *traefikv1alpha1.ErrorPage
	}
	tests := []struct {
		name    string
		args    args
		want    *dynamic.ErrorPage
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

			p := &Provider{}
			got, got1, err := p.createErrorPageMiddleware(tt.args.client, tt.args.namespace, tt.args.errorPage)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.createErrorPageMiddleware() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Provider.createErrorPageMiddleware() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Provider.createErrorPageMiddleware() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
