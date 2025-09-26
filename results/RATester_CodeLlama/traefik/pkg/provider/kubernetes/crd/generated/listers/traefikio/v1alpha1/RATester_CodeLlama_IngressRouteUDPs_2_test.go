package v1alpha1

import (
	"fmt"
	"testing"
)

func TestIngressRouteUDPs_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &ingressRouteUDPLister{}
	namespace := "namespace"
	s.IngressRouteUDPs(namespace)
}
