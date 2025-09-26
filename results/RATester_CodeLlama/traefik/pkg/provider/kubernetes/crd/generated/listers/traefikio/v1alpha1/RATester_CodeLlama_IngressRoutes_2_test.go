package v1alpha1

import (
	"fmt"
	"testing"
)

func TestIngressRoutes_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &ingressRouteLister{}
	namespace := "namespace"
	l := s.IngressRoutes(namespace)
	if l == nil {
		t.Error("expected non-nil")
	}
}
