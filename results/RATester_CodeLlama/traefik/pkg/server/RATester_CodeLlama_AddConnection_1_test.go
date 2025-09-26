package server

import (
	"fmt"
	"net"
	"testing"
)

func TestAddConnection_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &connectionTracker{}
	conn := &net.TCPConn{}

	c.AddConnection(conn)

	if c.isEmpty() {
		t.Error("connectionTracker should not be empty")
	}
}
