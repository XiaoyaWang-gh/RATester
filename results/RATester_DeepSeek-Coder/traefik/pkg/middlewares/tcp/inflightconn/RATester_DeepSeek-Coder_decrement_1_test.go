package inflightconn

import (
	"fmt"
	"testing"
)

func TestDecrement_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	i := &inFlightConn{
		connections: make(map[string]int64),
	}

	ip := "192.0.2.0"

	i.connections[ip] = 2

	i.decrement(ip)

	if i.connections[ip] != 1 {
		t.Errorf("Expected 1, got %d", i.connections[ip])
	}

	i.decrement(ip)

	if i.connections[ip] != 0 {
		t.Errorf("Expected 0, got %d", i.connections[ip])
	}

	i.decrement(ip)

	if i.connections[ip] != 0 {
		t.Errorf("Expected 0, got %d", i.connections[ip])
	}
}
