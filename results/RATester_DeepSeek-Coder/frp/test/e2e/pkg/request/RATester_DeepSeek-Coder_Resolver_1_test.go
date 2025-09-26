package request

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestResolver_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, "8.8.8.8:53")
		},
	}
	r.Resolver(resolver)
	if r.resolver != resolver {
		t.Errorf("Expected resolver to be %v, got %v", resolver, r.resolver)
	}
}
