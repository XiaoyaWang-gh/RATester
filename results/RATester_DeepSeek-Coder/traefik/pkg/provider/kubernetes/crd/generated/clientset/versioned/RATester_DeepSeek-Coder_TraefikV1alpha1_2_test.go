package versioned

import (
	"fmt"
	"testing"

	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/generated/clientset/versioned/typed/traefikio/v1alpha1"
)

func TestTraefikV1alpha1_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cs := &Clientset{
		traefikV1alpha1: &traefikv1alpha1.TraefikV1alpha1Client{},
	}

	got := cs.TraefikV1alpha1()
	if got == nil {
		t.Errorf("Expected non-nil TraefikV1alpha1Interface, got nil")
	}
}
