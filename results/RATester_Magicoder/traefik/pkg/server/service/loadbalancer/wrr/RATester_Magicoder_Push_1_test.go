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

	b := &Balancer{
		handlerMap: make(map[string]*namedHandler),
		handlers:   make([]*namedHandler, 0),
	}

	h := &namedHandler{
		name:   "test",
		weight: 1.0,
	}

	b.Push(h)

	if len(b.handlers) != 1 {
		t.Errorf("Expected handlers length to be 1, got %d", len(b.handlers))
	}

	if _, ok := b.handlerMap[h.name]; !ok {
		t.Errorf("Expected handlerMap to contain handler with name %s", h.name)
	}
}
