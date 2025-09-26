package wrr

import (
	"fmt"
	"testing"
)

func TestPop_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &Balancer{
		handlers: []*namedHandler{
			{name: "handler1"},
			{name: "handler2"},
			{name: "handler3"},
		},
	}

	popped := b.Pop()

	if popped.(*namedHandler).name != "handler3" {
		t.Errorf("Expected popped handler to be 'handler3', got %v", popped.(*namedHandler).name)
	}

	if len(b.handlers) != 2 {
		t.Errorf("Expected length of handlers to be 2, got %v", len(b.handlers))
	}
}
