package crd

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
)

func TestCreateDigestAuthMiddleware_1(t *testing.T) {
	type args struct {
		client     Client
		namespace  string
		digestAuth *traefikv1alpha1.DigestAuth
	}
	tests := []struct {
		name    string
		args    args
		want    *dynamic.DigestAuth
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

			got, err := createDigestAuthMiddleware(tt.args.client, tt.args.namespace, tt.args.digestAuth)
			if (err != nil) != tt.wantErr {
				t.Errorf("createDigestAuthMiddleware() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createDigestAuthMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
