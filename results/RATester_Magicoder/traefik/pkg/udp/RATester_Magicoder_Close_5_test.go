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

	err := conn.Close()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if _, ok := listener.conns[conn.rAddr.String()]; ok {
		t.Errorf("Connection was not removed from the listener's map")
	}
}
