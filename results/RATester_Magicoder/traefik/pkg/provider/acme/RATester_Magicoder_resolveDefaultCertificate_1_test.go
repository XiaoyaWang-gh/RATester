package acme

import (
	"context"
	"fmt"
	"testing"
)

func TestResolveDefaultCertificate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	provider := &Provider{
		Configuration: &Configuration{
			Email: "test@example.com",
		},
		resolvingDomains: make(map[string]struct{}),
	}

	domains := []string{"example.com", "www.example.com"}
	cert, err := provider.resolveDefaultCertificate(ctx, domains)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if cert == nil {
		t.Error("Expected a certificate, but got nil")
	}

	if len(cert.Certificate) == 0 {
		t.Error("Expected a non-empty certificate, but got an empty one")
	}

	if len(cert.PrivateKey) == 0 {
		t.Error("Expected a non-empty private key, but got an empty one")
	}
}
