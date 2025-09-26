package v1alpha1

import (
	"fmt"
	"reflect"
	"testing"
	time "time"

	versioned "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/generated/clientset/versioned"
	cache "k8s.io/client-go/tools/cache"
)

func TestNewServersTransportTCPInformer_1(t *testing.T) {
	type args struct {
		client       versioned.Interface
		namespace    string
		resyncPeriod time.Duration
		indexers     cache.Indexers
	}
	tests := []struct {
		name string
		args args
		want cache.SharedIndexInformer
	}{
		{
			name: "test_case_1",
			args: args{
				client:       nil,
				namespace:    "",
				resyncPeriod: 0,
				indexers:     nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewServersTransportTCPInformer(tt.args.client, tt.args.namespace, tt.args.resyncPeriod, tt.args.indexers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServersTransportTCPInformer() = %v, want %v", got, tt.want)
			}
		})
	}
}
