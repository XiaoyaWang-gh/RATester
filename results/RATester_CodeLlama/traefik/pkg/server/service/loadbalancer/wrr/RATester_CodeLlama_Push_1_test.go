package wrr

import (
	"fmt"
	"testing"
)

func TestPush_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &Balancer{}
	h := &namedHandler{}
	b.Push(h)
	if len(b.handlers) != 1 {
		t.Errorf("expected 1 handler, got %d", len(b.handlers))
	}
	if b.handlers[0] != h {
		t.Errorf("expected handler %v, got %v", h, b.handlers[0])
	}
}
