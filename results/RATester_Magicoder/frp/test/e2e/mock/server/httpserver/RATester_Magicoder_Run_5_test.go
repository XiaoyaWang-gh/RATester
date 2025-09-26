package httpserver

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestRun_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := &Server{
		bindAddr: "127.0.0.1",
		bindPort: 8080,
		handler:  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
	}

	err := server.Run()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if server.hs == nil {
		t.Fatal("Server is not initialized")
	}

	if server.hs.Addr != "127.0.0.1:8080" {
		t.Fatalf("Unexpected server address: %s", server.hs.Addr)
	}

	if server.hs.Handler == nil {
		t.Fatal("Server handler is not set")
	}

	if server.hs.TLSConfig != nil {
		t.Fatal("TLSConfig is not expected to be set")
	}

	if server.hs.ReadHeaderTimeout != time.Minute {
		t.Fatalf("Unexpected ReadHeaderTimeout: %v", server.hs.ReadHeaderTimeout)
	}

	server.Close()
}
