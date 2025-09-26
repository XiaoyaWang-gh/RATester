package request

import (
	"context"
	"fmt"
	"net"
	"testing"
)

func TestResolver_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, network, "8.8.8.8:53")
		},
	}

	req := &Request{}
	req.Resolver(resolver)

	if req.resolver != resolver {
		t.Errorf("Expected resolver to be set, but it was not.")
	}
}
