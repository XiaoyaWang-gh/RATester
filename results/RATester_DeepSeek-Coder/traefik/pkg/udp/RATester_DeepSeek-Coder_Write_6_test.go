package udp

import (
	"bytes"
	"fmt"
	"net"
	"testing"
	"time"
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
	listener.conns[conn.rAddr.String()] = conn

	testData := []byte("test data")
	n, err := conn.Write(testData)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if n != len(testData) {
		t.Errorf("Expected to write %d bytes, wrote %d", len(testData), n)
	}

	select {
	case data := <-conn.receiveCh:
		if !bytes.Equal(data, testData) {
			t.Errorf("Expected to receive %v, got %v", testData, data)
		}
	case <-time.After(time.Second):
		t.Errorf("Expected to receive data, timed out")
	}
}
