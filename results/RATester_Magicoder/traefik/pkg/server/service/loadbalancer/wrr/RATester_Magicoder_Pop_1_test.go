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

	poppedHandler := b.Pop()

	if poppedHandler.(*namedHandler).name != "handler3" {
		t.Errorf("Expected popped handler to be 'handler3', but got %s", poppedHandler.(*namedHandler).name)
	}

	if len(b.handlers) != 2 {
		t.Errorf("Expected handlers length to be 2, but got %d", len(b.handlers))
	}
}
