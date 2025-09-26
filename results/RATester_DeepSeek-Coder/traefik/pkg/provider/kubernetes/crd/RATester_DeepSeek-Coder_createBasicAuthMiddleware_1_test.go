package crd

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
)

func TestCreateBasicAuthMiddleware_1(t *testing.T) {
	type args struct {
		client    Client
		namespace string
		basicAuth *traefikv1alpha1.BasicAuth
	}
	tests := []struct {
		name    string
		args    args
		want    *dynamic.BasicAuth
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

			got, err := createBasicAuthMiddleware(tt.args.client, tt.args.namespace, tt.args.basicAuth)
			if (err != nil) != tt.wantErr {
				t.Errorf("createBasicAuthMiddleware() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createBasicAuthMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
