package inflightconn

import (
	"fmt"
	"testing"
)

func TestIncrement_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	i := &inFlightConn{
		name:           "test",
		maxConnections: 10,
		connections:    make(map[string]int64),
	}

	ip := "192.0.2.0"

	// Test initial increment
	err := i.increment(ip)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if i.connections[ip] != 1 {
		t.Errorf("Expected 1 connection, got %d", i.connections[ip])
	}

	// Test incrementing beyond maxConnections
	for j := 0; j < 10; j++ {
		err = i.increment(ip)
	}
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if i.connections[ip] != 10 {
		t.Errorf("Expected 10 connections, got %d", i.connections[ip])
	}
}
