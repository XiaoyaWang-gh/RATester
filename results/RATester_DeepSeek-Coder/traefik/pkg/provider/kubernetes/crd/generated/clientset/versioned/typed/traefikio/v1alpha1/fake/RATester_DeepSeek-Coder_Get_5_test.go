package fake

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	v1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGet_5(t *testing.T) {
	type args struct {
		ctx     context.Context
		name    string
		options v1.GetOptions
	}
	tests := []struct {
		name    string
		args    args
		want    *v1alpha1.TraefikService
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

			c := &FakeTraefikServices{}
			got, err := c.Get(tt.args.ctx, tt.args.name, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("FakeTraefikServices.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FakeTraefikServices.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
