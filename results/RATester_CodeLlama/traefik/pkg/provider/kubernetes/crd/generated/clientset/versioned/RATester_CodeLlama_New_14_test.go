package versioned

import (
	"fmt"
	"testing"

	rest "k8s.io/client-go/rest"
)

func TestNew_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var c rest.Interface
	cs := New(c)
	if cs.traefikV1alpha1 == nil {
		t.Error("expected a non-nil traefikV1alpha1 client")
	}
	if cs.DiscoveryClient == nil {
		t.Error("expected a non-nil DiscoveryClient")
	}
}
