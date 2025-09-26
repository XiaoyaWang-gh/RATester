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

	// Test incrementing a connection
	err := i.increment("192.168.1.1")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if i.connections["192.168.1.1"] != 1 {
		t.Errorf("Expected 1 connection, got %d", i.connections["192.168.1.1"])
	}

	// Test incrementing beyond max connections
	err = i.increment("192.168.1.1")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if i.connections["192.168.1.1"] != 1 {
		t.Errorf("Expected 1 connection, got %d", i.connections["192.168.1.1"])
	}
}
