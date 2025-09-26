package udp

import (
	"fmt"
	"net"
	"testing"
)

func TestWrite_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	listener := &Listener{
		pConn: &net.UDPConn{},
		conns: make(map[string]*Conn),
	}
	conn := &Conn{
		listener:  listener,
		rAddr:     &net.UDPAddr{},
		receiveCh: make(chan []byte),
		readCh:    make(chan []byte),
		sizeCh:    make(chan int),
		msgs:      make([][]byte, 0),
		doneCh:    make(chan struct{}),
	}

	data := []byte("test data")
	n, err := conn.Write(data)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if n != len(data) {
		t.Errorf("Expected to write %d bytes, but wrote %d", len(data), n)
	}
}
