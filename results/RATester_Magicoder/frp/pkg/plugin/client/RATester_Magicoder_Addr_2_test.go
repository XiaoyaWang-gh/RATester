package client

import (
	"fmt"
	"net"
	"testing"
)

func TestAddr_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &Listener{
		conns:  make(chan net.Conn),
		closed: false,
	}

	addr := l.Addr()

	if addr != (*net.TCPAddr)(nil) {
		t.Errorf("Expected nil, got %v", addr)
	}
}
