package net

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestNewWebsocketListener_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	wl := NewWebsocketListener(ln)
	defer wl.Close()

	conn, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	select {
	case <-wl.acceptCh:
		// Connection accepted
	case <-time.After(time.Second):
		t.Fatal("Timeout waiting for connection")
	}
}
