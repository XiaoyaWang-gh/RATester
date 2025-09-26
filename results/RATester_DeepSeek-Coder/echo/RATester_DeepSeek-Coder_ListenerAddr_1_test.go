package echo

import (
	"fmt"
	"net"
	"testing"
)

func TestListenerAddr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Echo{}

	// Test when Listener is nil
	addr := e.ListenerAddr()
	if addr != nil {
		t.Errorf("Expected nil, got %v", addr)
	}

	// Test when Listener is not nil
	e.Listener = &net.TCPListener{}
	addr = e.ListenerAddr()
	if addr == nil {
		t.Errorf("Expected non-nil address")
	}
}
