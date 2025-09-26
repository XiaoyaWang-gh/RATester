package helpers

import (
	"fmt"
	"testing"
)

func TestTCPListen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l, addr, err := TCPListen()
	if err != nil {
		t.Fatalf("TCPListen() error = %v", err)
	}
	defer l.Close()

	if addr == nil {
		t.Fatalf("TCPListen() addr = nil")
	}

	if addr.Port == 0 {
		t.Fatalf("TCPListen() addr.Port = 0")
	}

	if addr.IP == nil {
		t.Fatalf("TCPListen() addr.IP = nil")
	}
}
