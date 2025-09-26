package echo

import (
	"fmt"
	"net"
	"testing"
)

func TestTLSListenerAddr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Echo{}

	// Test when TLSListener is nil
	e.startupMutex.Lock()
	e.TLSListener = nil
	e.startupMutex.Unlock()
	addr := e.TLSListenerAddr()
	if addr != nil {
		t.Errorf("Expected nil, got %v", addr)
	}

	// Test when TLSListener is not nil
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	e.startupMutex.Lock()
	e.TLSListener = l
	e.startupMutex.Unlock()
	addr = e.TLSListenerAddr()
	if addr == nil {
		t.Errorf("Expected non-nil address, got nil")
	} else {
		expected := l.Addr().String()
		actual := addr.String()
		if expected != actual {
			t.Errorf("Expected %q, got %q", expected, actual)
		}
	}
}
