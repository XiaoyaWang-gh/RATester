package fake

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	v1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestUpdate_1(t *testing.T) {
	type args struct {
		ctx            context.Context
		traefikService *v1alpha1.TraefikService
		opts           v1.UpdateOptions
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

			c := &FakeTraefikServices{
				Fake: &FakeTraefikV1alpha1{},
			}
			got, err := c.Update(tt.args.ctx, tt.args.traefikService, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("FakeTraefikServices.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FakeTraefikServices.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
