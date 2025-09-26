package server

import (
	"fmt"
	"net"
	"testing"
)

func TestRemoveConnection_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// CONTEXT_0
	c := &connectionTracker{}

	// CONTEXT_1
	c.conns = make(map[net.Conn]struct{})

	// CONTEXT_2
	conn := &net.TCPConn{}

	// CONTEXT_3
	c.AddConnection(conn)

	// CONTEXT_4
	c.RemoveConnection(conn)

	// CONTEXT_5
	if _, ok := c.conns[conn]; ok {
		t.Errorf("RemoveConnection() failed, connection still in the tracked connections list")
	}
}
