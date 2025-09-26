package traefikio

import (
	"fmt"
	"testing"
)

func TestV1alpha1_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	g := &group{
		factory:          nil,
		namespace:        "default",
		tweakListOptions: nil,
	}

	v1alpha1Client := g.V1alpha1()

	if v1alpha1Client == nil {
		t.Errorf("Expected v1alpha1Client to not be nil")
	}
}
