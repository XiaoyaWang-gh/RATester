package inflightconn

import (
	"fmt"
	"sync"
	"testing"
)

func TestDecrement_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	i := &inFlightConn{
		name:           "test",
		next:           nil,
		maxConnections: 10,
		mu:             sync.Mutex{},
		connections:    map[string]int64{"127.0.0.1": 1},
	}

	ip := "127.0.0.1"

	i.decrement(ip)

	if i.connections[ip] != 0 {
		t.Errorf("decrement failed, want 0, got %d", i.connections[ip])
	}
}
