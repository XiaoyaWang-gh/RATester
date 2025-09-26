package integration

import (
	"fmt"
	"testing"
)

func TestStartFakeDNSServer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	traefikIP := "127.0.0.1"
	srv := startFakeDNSServer(traefikIP)

	if srv == nil {
		t.Errorf("Expected a non-nil server, got nil")
	}

	if srv.Addr != ":5053" {
		t.Errorf("Expected server address to be ':5053', got '%s'", srv.Addr)
	}

	if srv.Net != "udp" {
		t.Errorf("Expected server network to be 'udp', got '%s'", srv.Net)
	}

	if srv.Handler == nil {
		t.Errorf("Expected a non-nil handler, got nil")
	}

	// TODO: Add more assertions to test the server's behavior
}
