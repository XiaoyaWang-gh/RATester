package wrr

import (
	"fmt"
	"testing"
)

func TestSwap_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &Balancer{
		handlers: []*namedHandler{
			{name: "handler1"},
			{name: "handler2"},
		},
	}

	b.Swap(0, 1)

	if b.handlers[0].name != "handler2" || b.handlers[1].name != "handler1" {
		t.Errorf("Swap failed. Expected handlers to be swapped. Got: %v", b.handlers)
	}
}
