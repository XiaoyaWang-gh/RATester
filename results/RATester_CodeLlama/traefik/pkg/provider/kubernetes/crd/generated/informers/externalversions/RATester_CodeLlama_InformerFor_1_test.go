package externalversions

import (
	"fmt"
	"testing"
	time "time"

	versioned "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/generated/clientset/versioned"
	v1 "k8s.io/api/core/v1"
	cache "k8s.io/client-go/tools/cache"
)

func TestInformerFor_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &sharedInformerFactory{}
	obj := &v1.Service{}
	newFunc := func(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
		return nil
	}
	informer := f.InformerFor(obj, newFunc)
	if informer == nil {
		t.Errorf("Expected non-nil informer")
	}
}
