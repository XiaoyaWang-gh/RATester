package pmux

import (
	"fmt"
	"net"
	"testing"
)

func TestNewPortListener_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	connCh := make(chan *PortConn)
	addr := &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8080}

	pl := NewPortListener(connCh, addr)

	if pl.connCh != connCh {
		t.Errorf("Expected connCh to be %v, but got %v", connCh, pl.connCh)
	}

	if pl.addr != addr {
		t.Errorf("Expected addr to be %v, but got %v", addr, pl.addr)
	}
}
