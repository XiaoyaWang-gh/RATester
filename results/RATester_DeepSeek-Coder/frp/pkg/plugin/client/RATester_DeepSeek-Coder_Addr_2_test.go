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
		conns: make(chan net.Conn),
	}

	expected := (*net.TCPAddr)(nil)
	actual := l.Addr()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
