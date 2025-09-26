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

	p := &Provider{}
	ctx := context.Background()
	domains := []string{"test.com"}
	tlsStore := "test"
	p.resolveDomains(ctx, domains, tlsStore)
}
