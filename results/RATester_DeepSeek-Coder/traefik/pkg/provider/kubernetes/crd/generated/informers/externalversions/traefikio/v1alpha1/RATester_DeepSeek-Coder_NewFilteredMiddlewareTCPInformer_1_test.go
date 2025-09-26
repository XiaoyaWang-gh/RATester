package v1alpha1

import (
	"fmt"
	"reflect"
	"testing"
	time "time"

	versioned "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/generated/clientset/versioned"
	internalinterfaces "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/generated/informers/externalversions/internalinterfaces"
	cache "k8s.io/client-go/tools/cache"
)

func TestNewFilteredMiddlewareTCPInformer_1(t *testing.T) {
	type args struct {
		client           versioned.Interface
		namespace        string
		resyncPeriod     time.Duration
		indexers         cache.Indexers
		tweakListOptions internalinterfaces.TweakListOptionsFunc
	}
	tests := []struct {
		name string
		args args
		want cache.SharedIndexInformer
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

			if got := NewFilteredMiddlewareTCPInformer(tt.args.client, tt.args.namespace, tt.args.resyncPeriod, tt.args.indexers, tt.args.tweakListOptions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilteredMiddlewareTCPInformer() = %v, want %v", got, tt.want)
			}
		})
	}
}
