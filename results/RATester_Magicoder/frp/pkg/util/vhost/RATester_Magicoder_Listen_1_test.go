package vhost

import (
	"context"
	"fmt"
	"testing"
)

func TestListen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	muxer := &Muxer{
		registryRouter: &Routers{
			indexByDomain: make(map[string]routerByHTTPUser),
		},
	}

	ctx := context.Background()
	cfg := &RouteConfig{
		Domain:          "example.com",
		Location:        "/",
		RouteByHTTPUser: "user",
	}

	l, err := muxer.Listen(ctx, cfg)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if l.name != cfg.Domain {
		t.Errorf("Expected name %s, got %s", cfg.Domain, l.name)
	}

	if l.location != cfg.Location {
		t.Errorf("Expected location %s, got %s", cfg.Location, l.location)
	}

	if l.routeByHTTPUser != cfg.RouteByHTTPUser {
		t.Errorf("Expected routeByHTTPUser %s, got %s", cfg.RouteByHTTPUser, l.routeByHTTPUser)
	}

	if l.mux != muxer {
		t.Errorf("Expected muxer %p, got %p", muxer, l.mux)
	}

	if l.ctx != ctx {
		t.Errorf("Expected context %p, got %p", ctx, l.ctx)
	}

	if _, ok := muxer.registryRouter.indexByDomain[cfg.Domain]; !ok {
		t.Errorf("Expected domain %s to be in registryRouter", cfg.Domain)
	}
}
