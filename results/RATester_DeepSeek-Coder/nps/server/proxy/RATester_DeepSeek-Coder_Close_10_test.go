package proxy

import (
	"fmt"
	"net"
	"testing"
)

func TestClose_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := &UdpModeServer{
		listener: &net.UDPConn{},
	}

	err := server.Close()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
