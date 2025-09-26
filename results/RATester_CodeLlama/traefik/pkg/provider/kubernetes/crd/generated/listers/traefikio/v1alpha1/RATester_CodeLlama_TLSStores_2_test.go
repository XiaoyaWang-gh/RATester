package v1alpha1

import (
	"fmt"
	"testing"
)

func TestTLSStores_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &tLSStoreLister{}
	namespace := "test"
	lister := s.TLSStores(namespace)
	if lister == nil {
		t.Errorf("TLSStores() returned nil")
	}
}
