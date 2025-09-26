package acme

import (
	"fmt"
	"testing"

	traefiktls "github.com/traefik/traefik/v3/pkg/tls"
)

func TestSetTLSManager_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := &traefiktls.Manager{}
	provider := &Provider{}

	provider.SetTLSManager(manager)

	if provider.tlsManager != manager {
		t.Errorf("Expected tlsManager to be %v, but got %v", manager, provider.tlsManager)
	}
}
