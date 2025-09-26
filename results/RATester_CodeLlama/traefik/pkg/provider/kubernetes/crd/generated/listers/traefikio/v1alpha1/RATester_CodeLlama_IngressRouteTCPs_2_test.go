package v1alpha1

import (
	"fmt"
	"testing"
)

func TestIngressRouteTCPs_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &ingressRouteTCPLister{}
	namespace := "namespace"
	l := s.IngressRouteTCPs(namespace)
	if l == nil {
		t.Error("IngressRouteTCPs() returned nil")
	}
}
