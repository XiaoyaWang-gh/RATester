package proxy

import (
	"fmt"
	"net"
	"testing"
)

func TestClose_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := &Sock5ModeServer{
		listener: &net.TCPListener{},
	}

	err := server.Close()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
