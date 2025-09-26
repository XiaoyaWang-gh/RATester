package v1alpha1

import (
	"fmt"
	"testing"
)

func TestServersTransportTCPs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := version{}
	if got := v.ServersTransportTCPs(); got == nil {
		t.Error("ServersTransportTCPs() = nil")
	}
}
