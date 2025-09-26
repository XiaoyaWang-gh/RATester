package v1alpha1

import (
	"fmt"
	"testing"

	rest "k8s.io/client-go/rest"
)

func TestNewForConfigOrDie_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &rest.Config{}
	client := NewForConfigOrDie(c)
	if client == nil {
		t.Errorf("expected a client, got nil")
	}
}
