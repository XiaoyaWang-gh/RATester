package v1alpha1

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	v1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
)

func TestPatch_2(t *testing.T) {
	type args struct {
		ctx          context.Context
		name         string
		pt           types.PatchType
		data         []byte
		opts         v1.PatchOptions
		subresources []string
	}
	tests := []struct {
		name    string
		args    args
		want    *v1alpha1.IngressRouteTCP
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

			c := &ingressRouteTCPs{}
			got, err := c.Patch(tt.args.ctx, tt.args.name, tt.args.pt, tt.args.data, tt.args.opts, tt.args.subresources...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Patch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Patch() got = %v, want %v", got, tt.want)
			}
		})
	}
}
