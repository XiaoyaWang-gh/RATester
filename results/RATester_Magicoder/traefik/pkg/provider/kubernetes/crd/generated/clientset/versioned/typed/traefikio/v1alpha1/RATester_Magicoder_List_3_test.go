package v1alpha1

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	v1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestList_3(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts v1.ListOptions
	}
	tests := []struct {
		name    string
		c       *ingressRouteTCPs
		args    args
		want    *v1alpha1.IngressRouteTCPList
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

			got, err := tt.c.List(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("ingressRouteTCPs.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ingressRouteTCPs.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
