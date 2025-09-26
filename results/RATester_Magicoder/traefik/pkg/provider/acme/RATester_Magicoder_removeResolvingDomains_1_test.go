package acme

import (
	"fmt"
	"testing"
)

func TestRemoveResolvingDomains_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &Provider{
		resolvingDomains: map[string]struct{}{
			"example.com": {},
			"test.com":    {},
		},
	}

	domains := []string{"example.com", "test.com"}
	provider.removeResolvingDomains(domains)

	if len(provider.resolvingDomains) != 0 {
		t.Errorf("Expected resolvingDomains to be empty, but got %d", len(provider.resolvingDomains))
	}
}
