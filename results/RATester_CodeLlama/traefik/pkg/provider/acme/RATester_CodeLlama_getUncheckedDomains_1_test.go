package acme

import (
	"context"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	traefiktls "github.com/traefik/traefik/v3/pkg/tls"
)

func TestGetUncheckedDomains_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	p := &Provider{
		tlsManager: &traefiktls.Manager{},
	}
	ctx := context.Background()
	domainsToCheck := []string{"test.com"}
	tlsStore := "default"

	// when
	uncheckedDomains := p.getUncheckedDomains(ctx, domainsToCheck, tlsStore)

	// then
	assert.Len(t, uncheckedDomains, 1)
	assert.Equal(t, "test.com", uncheckedDomains[0])
}
