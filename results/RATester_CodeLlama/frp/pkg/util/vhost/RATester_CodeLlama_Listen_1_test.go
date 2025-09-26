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

	ctx := context.Background()
	cfg := &RouteConfig{
		Domain:          "test.com",
		Location:        "/",
		RewriteHost:     "test.com",
		Username:        "test",
		Password:        "test",
		RouteByHTTPUser: "test",
	}
	muxer := &Muxer{}
	l, err := muxer.Listen(ctx, cfg)
	if err != nil {
		t.Errorf("Listen() error = %v, want nil", err)
		return
	}
	if l.name != cfg.Domain {
		t.Errorf("Listen() name = %v, want %v", l.name, cfg.Domain)
	}
	if l.location != cfg.Location {
		t.Errorf("Listen() location = %v, want %v", l.location, cfg.Location)
	}
	if l.routeByHTTPUser != cfg.RouteByHTTPUser {
		t.Errorf("Listen() routeByHTTPUser = %v, want %v", l.routeByHTTPUser, cfg.RouteByHTTPUser)
	}
	if l.rewriteHost != cfg.RewriteHost {
		t.Errorf("Listen() rewriteHost = %v, want %v", l.rewriteHost, cfg.RewriteHost)
	}
	if l.username != cfg.Username {
		t.Errorf("Listen() username = %v, want %v", l.username, cfg.Username)
	}
	if l.password != cfg.Password {
		t.Errorf("Listen() password = %v, want %v", l.password, cfg.Password)
	}
	if l.mux != muxer {
		t.Errorf("Listen() mux = %v, want %v", l.mux, muxer)
	}
}
