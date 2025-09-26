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

	ip := "127.0.0.1"

	err := i.increment(ip)
	if err != nil {
		t.Errorf("increment() error = %v", err)
		return
	}

	if i.connections[ip] != 1 {
		t.Errorf("increment() connections[%s] = %d, want %d", ip, i.connections[ip], 1)
	}

	err = i.increment(ip)
	if err != nil {
		t.Errorf("increment() error = %v", err)
		return
	}

	if i.connections[ip] != 2 {
		t.Errorf("increment() connections[%s] = %d, want %d", ip, i.connections[ip], 2)
	}

	err = i.increment(ip)
	if err != nil {
		t.Errorf("increment() error = %v", err)
		return
	}

	if i.connections[ip] != 3 {
		t.Errorf("increment() connections[%s] = %d, want %d", ip, i.connections[ip], 3)
	}

	err = i.increment(ip)
	if err == nil {
		t.Errorf("increment() error = nil, want %v", fmt.Errorf("max number of connections reached for %s", ip))
		return
	}

	if i.connections[ip] != 3 {
		t.Errorf("increment() connections[%s] = %d, want %d", ip, i.connections[ip], 3)
	}
}
