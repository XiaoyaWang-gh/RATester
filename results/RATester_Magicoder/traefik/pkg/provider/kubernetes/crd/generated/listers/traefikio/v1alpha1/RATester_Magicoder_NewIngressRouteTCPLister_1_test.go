package v1alpha1

import (
	"fmt"
	"testing"

	"k8s.io/client-go/tools/cache"
)

func TestNewIngressRouteTCPLister_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	indexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{})
	lister := NewIngressRouteTCPLister(indexer)

	// Test case 1: Test if the lister is not nil
	if lister == nil {
		t.Errorf("Expected lister to be not nil")
	}

	// Test case 2: Test if the lister implements the correct interface
	_, ok := lister.(IngressRouteTCPLister)
	if !ok {
		t.Errorf("Expected lister to implement IngressRouteTCPLister")
	}

	// Test case 3: Test if the lister returns the correct indexer
	if lister.(*ingressRouteTCPLister).indexer != indexer {
		t.Errorf("Expected lister to return the correct indexer")
	}
}
