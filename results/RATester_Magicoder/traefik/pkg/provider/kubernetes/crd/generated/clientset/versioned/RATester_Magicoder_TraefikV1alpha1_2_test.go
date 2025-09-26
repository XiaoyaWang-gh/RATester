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

	client := &Clientset{
		traefikV1alpha1: &traefikv1alpha1.TraefikV1alpha1Client{},
	}

	traefikV1alpha1Client := client.TraefikV1alpha1()

	if traefikV1alpha1Client == nil {
		t.Error("Expected TraefikV1alpha1Client, but got nil")
	}
}
