package acme

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestGetUncheckedDomains_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	domainsToCheck := []string{"example.com", "test.com"}
	tlsStore := "testStore"

	provider := &Provider{
		certificates: []*CertAndStore{
			{
				Certificate: Certificate{
					Domain: types.Domain{Main: "example.com"},
				},
			},
		},
		resolvingDomains: make(map[string]struct{}),
	}

	uncheckedDomains := provider.getUncheckedDomains(ctx, domainsToCheck, tlsStore)

	if len(uncheckedDomains) != 1 {
		t.Errorf("Expected 1 unchecked domain, got %d", len(uncheckedDomains))
	}

	if uncheckedDomains[0] != "test.com" {
		t.Errorf("Expected unchecked domain to be 'test.com', got %s", uncheckedDomains[0])
	}
}
