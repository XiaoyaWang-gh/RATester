package v1alpha1

import (
	"fmt"
	"reflect"
	"testing"

	"k8s.io/client-go/tools/cache"
)

func TestMiddlewareTCPs_2(t *testing.T) {
	type fields struct {
		indexer cache.Indexer
	}
	type args struct {
		namespace string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   MiddlewareTCPNamespaceLister
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

			s := &middlewareTCPLister{
				indexer: tt.fields.indexer,
			}
			if got := s.MiddlewareTCPs(tt.args.namespace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middlewareTCPLister.MiddlewareTCPs() = %v, want %v", got, tt.want)
			}
		})
	}
}
