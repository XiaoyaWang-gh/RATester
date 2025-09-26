package httpserver

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"
)

func TestRun_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Server{
		bindAddr:  "127.0.0.1",
		bindPort:  8080,
		handler:   http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
		tlsConfig: &tls.Config{},
	}

	err := s.Run()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if s.hs == nil {
		t.Fatalf("http server not initialized")
	}

	if s.hs.Addr != "127.0.0.1:8080" {
		t.Fatalf("unexpected addr, want: 127.0.0.1:8080, got: %s", s.hs.Addr)
	}

	if s.hs.Handler != s.handler {
		t.Fatalf("unexpected handler, want: %v, got: %v", s.handler, s.hs.Handler)
	}

	if s.hs.TLSConfig != s.tlsConfig {
		t.Fatalf("unexpected tls config, want: %v, got: %v", s.tlsConfig, s.hs.TLSConfig)
	}

	s.Close()
}
