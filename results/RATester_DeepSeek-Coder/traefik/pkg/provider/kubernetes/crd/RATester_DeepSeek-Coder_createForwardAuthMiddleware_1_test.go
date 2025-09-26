package crd

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
)

func TestCreateForwardAuthMiddleware_1(t *testing.T) {
	type args struct {
		k8sClient Client
		namespace string
		auth      *traefikv1alpha1.ForwardAuth
	}
	tests := []struct {
		name    string
		args    args
		want    *dynamic.ForwardAuth
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

			got, err := createForwardAuthMiddleware(tt.args.k8sClient, tt.args.namespace, tt.args.auth)
			if (err != nil) != tt.wantErr {
				t.Errorf("createForwardAuthMiddleware() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createForwardAuthMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
