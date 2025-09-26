package v1alpha1

import (
	"fmt"
	"testing"
)

func TestTLSStores_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &version{}
	if got := v.TLSStores(); got == nil {
		t.Error("expected non-nil")
	}
}
