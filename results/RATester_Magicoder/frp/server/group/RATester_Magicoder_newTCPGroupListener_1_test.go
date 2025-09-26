package group

import (
	"fmt"
	"net"
	"testing"
)

func TestNewTCPGroupListener_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	name := "test"
	group := &TCPGroup{}
	addr := &net.TCPAddr{}

	listener := newTCPGroupListener(name, group, addr)

	if listener.groupName != name {
		t.Errorf("Expected groupName to be %s, but got %s", name, listener.groupName)
	}

	if listener.group != group {
		t.Errorf("Expected group to be %v, but got %v", group, listener.group)
	}

	if listener.addr != addr {
		t.Errorf("Expected addr to be %v, but got %v", addr, listener.addr)
	}

	if len(listener.closeCh) != 0 {
		t.Errorf("Expected closeCh to be empty, but got %v", listener.closeCh)
	}
}
