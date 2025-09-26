package wrr

import (
	"fmt"
	"testing"
)

func TestNextServer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &Balancer{
		handlers: []*namedHandler{
			{name: "server1", weight: 1},
			{name: "server2", weight: 2},
			{name: "server3", weight: 3},
		},
		status: map[string]struct{}{
			"server1": {},
			"server2": {},
			"server3": {},
		},
	}

	handler, err := b.nextServer()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if handler == nil {
		t.Fatalf("Expected handler, got nil")
	}

	if handler.name != "server1" && handler.name != "server2" && handler.name != "server3" {
		t.Fatalf("Expected server1, server2, or server3, got %s", handler.name)
	}
}
