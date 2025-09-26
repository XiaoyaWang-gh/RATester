package v1alpha1

import (
	"fmt"
	"reflect"
	"testing"
	time "time"

	versioned "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/generated/clientset/versioned"
	cache "k8s.io/client-go/tools/cache"
)

func TestDefaultInformer_1(t *testing.T) {
	type args struct {
		client       versioned.Interface
		resyncPeriod time.Duration
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

			f := &traefikServiceInformer{}
			if got := f.defaultInformer(tt.args.client, tt.args.resyncPeriod); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultInformer() = %v, want %v", got, tt.want)
			}
		})
	}
}
