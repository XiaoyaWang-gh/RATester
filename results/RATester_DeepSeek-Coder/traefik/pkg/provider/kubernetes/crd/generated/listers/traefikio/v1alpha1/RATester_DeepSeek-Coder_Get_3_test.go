package v1alpha1

import (
	"fmt"
	"reflect"
	"testing"

	v1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	"k8s.io/client-go/tools/cache"
)

func TestGet_3(t *testing.T) {
	type fields struct {
		indexer   cache.Indexer
		namespace string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
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

			s := traefikServiceNamespaceLister{
				indexer:   tt.fields.indexer,
				namespace: tt.fields.namespace,
			}
			got, err := s.Get(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("traefikServiceNamespaceLister.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("traefikServiceNamespaceLister.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
