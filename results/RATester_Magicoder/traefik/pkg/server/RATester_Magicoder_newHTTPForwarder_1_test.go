package server

import (
	"fmt"
	"net"
	"testing"
)

func TestNewHTTPForwarder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	forwarder := newHTTPForwarder(ln)

	if forwarder.Listener != ln {
		t.Errorf("Expected forwarder.Listener to be %v, but got %v", ln, forwarder.Listener)
	}

	if cap(forwarder.connChan) != 0 {
		t.Errorf("Expected forwarder.connChan to have a capacity of 0, but got %d", cap(forwarder.connChan))
	}

	if cap(forwarder.errChan) != 0 {
		t.Errorf("Expected forwarder.errChan to have a capacity of 0, but got %d", cap(forwarder.errChan))
	}
}
