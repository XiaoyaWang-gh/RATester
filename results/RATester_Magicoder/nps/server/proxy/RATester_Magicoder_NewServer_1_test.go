package proxy

import (
	"fmt"
	"testing"
)

func TestNewServer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := &httpServer{
		httpPort: 8080,
	}

	httpServer := server.NewServer(server.httpPort, "http")

	if httpServer == nil {
		t.Error("Expected a non-nil http.Server, but got nil")
	}

	if httpServer.Addr != ":8080" {
		t.Errorf("Expected http.Server.Addr to be ':8080', but got '%s'", httpServer.Addr)
	}
}
