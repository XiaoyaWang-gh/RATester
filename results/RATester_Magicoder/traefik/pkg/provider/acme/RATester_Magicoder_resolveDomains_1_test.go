package acme

import (
	"context"
	"fmt"
	"testing"
)

func TestResolveDomains_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	provider := &Provider{
		ResolverName: "test",
	}

	domains := []string{"example.com", "www.example.com"}
	tlsStore := "/path/to/tls/store"

	provider.resolveDomains(ctx, domains, tlsStore)

	// Add assertions here to verify the behavior of the method under test
}
