package externalversions

import (
	"fmt"
	reflect "reflect"
	"testing"

	cache "k8s.io/client-go/tools/cache"
)

func TestWithTransform_1(t *testing.T) {
	type args struct {
		transform cache.TransformFunc
	}
	tests := []struct {
		name string
		args args
		want SharedInformerOption
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

			if got := WithTransform(tt.args.transform); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}
