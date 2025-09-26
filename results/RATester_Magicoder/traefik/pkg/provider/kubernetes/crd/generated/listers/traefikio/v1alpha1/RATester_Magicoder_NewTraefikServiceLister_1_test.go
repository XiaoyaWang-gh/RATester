package v1alpha1

import (
	"fmt"
	"testing"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

func TestNewTraefikServiceLister_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	indexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{})
	lister := NewTraefikServiceLister(indexer)

	// Test case 1: Test if the lister is not nil
	if lister == nil {
		t.Errorf("NewTraefikServiceLister() returned nil")
	}

	// Test case 2: Test if the lister implements the correct interface
	_, ok := lister.(TraefikServiceLister)
	if !ok {
		t.Errorf("NewTraefikServiceLister() does not implement TraefikServiceLister")
	}

	// Test case 3: Test if the lister can list items
	_, err := lister.List(labels.Everything())
	if err != nil {
		t.Errorf("NewTraefikServiceLister() failed to list items: %v", err)
	}

	// Test case 4: Test if the lister can list items in a namespace
	_, err = lister.TraefikServices("default").List(labels.Everything())
	if err != nil {
		t.Errorf("NewTraefikServiceLister() failed to list items in a namespace: %v", err)
	}
}
