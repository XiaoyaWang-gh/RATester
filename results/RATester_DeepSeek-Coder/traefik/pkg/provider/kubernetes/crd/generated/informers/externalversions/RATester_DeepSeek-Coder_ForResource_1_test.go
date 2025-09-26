package externalversions

import (
	"fmt"
	reflect "reflect"
	"testing"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

func TestForResource_1(t *testing.T) {
	type args struct {
		resource schema.GroupVersionResource
	}
	tests := []struct {
		name    string
		f       *sharedInformerFactory
		args    args
		want    GenericInformer
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

			got, err := tt.f.ForResource(tt.args.resource)
			if (err != nil) != tt.wantErr {
				t.Errorf("sharedInformerFactory.ForResource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sharedInformerFactory.ForResource() = %v, want %v", got, tt.want)
			}
		})
	}
}
