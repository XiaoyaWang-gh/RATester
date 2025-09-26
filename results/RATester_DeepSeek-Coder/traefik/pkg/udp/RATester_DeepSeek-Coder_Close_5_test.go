package udp

import (
	"fmt"
	"net"
	"testing"
)

func TestClose_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	listener := &Listener{
		conns: make(map[string]*Conn),
	}
	conn := &Conn{
		listener: listener,
		rAddr:    &net.UDPAddr{},
	}
	listener.conns[conn.rAddr.String()] = conn

	conn.Close()

	if _, ok := listener.conns[conn.rAddr.String()]; ok {
		t.Errorf("Expected connection to be removed from listener's connections map")
	}
}
