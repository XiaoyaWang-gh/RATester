package fake

import (
	"fmt"
	"testing"

	fakediscovery "k8s.io/client-go/discovery/fake"
)

func TestDiscovery_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cs := &Clientset{
		discovery: &fakediscovery.FakeDiscovery{},
	}

	discovery := cs.Discovery()
	if discovery == nil {
		t.Errorf("Expected non-nil discovery client, got nil")
	}
}
