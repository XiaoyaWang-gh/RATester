package udp

import (
	"fmt"
	"net"
	"testing"
)

func TestAddr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	listener := &Listener{
		pConn: &net.UDPConn{},
	}

	expected := listener.pConn.LocalAddr()
	actual := listener.Addr()

	if expected != actual {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
