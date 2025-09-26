package v1alpha1

import (
	"fmt"
	"reflect"
	"testing"

	"k8s.io/client-go/tools/cache"
)

func TestNewIngressRouteLister_1(t *testing.T) {
	type args struct {
		indexer cache.Indexer
	}
	tests := []struct {
		name string
		args args
		want IngressRouteLister
	}{
		{
			name: "test",
			args: args{
				indexer: cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{}),
			},
			want: &ingressRouteLister{indexer: cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{})},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewIngressRouteLister(tt.args.indexer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIngressRouteLister() = %v, want %v", got, tt.want)
			}
		})
	}
}
