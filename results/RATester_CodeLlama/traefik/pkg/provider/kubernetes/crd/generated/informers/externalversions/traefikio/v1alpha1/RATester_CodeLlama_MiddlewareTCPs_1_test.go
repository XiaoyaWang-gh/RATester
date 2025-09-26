package v1alpha1

import (
	"fmt"
	"testing"
)

func TestMiddlewareTCPs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &version{}
	if got := v.MiddlewareTCPs(); got == nil {
		t.Error("expected non-nil MiddlewareTCPInformer")
	}
}
