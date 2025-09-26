package tailscale

import (
	"context"
	"fmt"
	"testing"

	traefiktls "github.com/traefik/traefik/v3/pkg/tls"
)

func TestFetchCerts_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	provider := &Provider{
		certByDomain: make(map[string]traefiktls.Certificate),
	}

	domains := []string{"example.com", "example.org"}
	provider.fetchCerts(ctx, domains)

	for _, domain := range domains {
		_, ok := provider.certByDomain[domain]
		if !ok {
			t.Errorf("Expected certificate for domain %q, but not found", domain)
		}
	}
}
