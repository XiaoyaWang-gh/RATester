package versioned

import (
	"fmt"
	"testing"

	discovery "k8s.io/client-go/discovery"
)

func TestDiscovery_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cs := &Clientset{
		DiscoveryClient: &discovery.DiscoveryClient{},
	}

	discoveryClient := cs.Discovery()
	if discoveryClient == nil {
		t.Fatal("Expected non-nil DiscoveryClient")
	}

	// Add more tests as needed
}
