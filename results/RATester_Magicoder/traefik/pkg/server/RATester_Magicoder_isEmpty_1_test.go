package server

import (
	"fmt"
	"net"
	"testing"
)

func TestIsEmpty_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ct := &connectionTracker{
		conns: make(map[net.Conn]struct{}),
	}

	// Test case 1: Empty map
	isEmpty := ct.isEmpty()
	if !isEmpty {
		t.Errorf("Expected isEmpty to be true, but got false")
	}

	// Test case 2: Non-empty map
	conn := &net.TCPConn{}
	ct.conns[conn] = struct{}{}
	isEmpty = ct.isEmpty()
	if isEmpty {
		t.Errorf("Expected isEmpty to be false, but got true")
	}
}
