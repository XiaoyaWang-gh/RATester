package v1alpha1

import (
	"fmt"
	"reflect"
	"testing"

	"k8s.io/client-go/tools/cache"
)

func TestIngressRoutes_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	indexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{})
	lister := &ingressRouteLister{indexer: indexer}

	namespace := "test-namespace"
	expected := ingressRouteNamespaceLister{indexer: indexer, namespace: namespace}

	result := lister.IngressRoutes(namespace)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
