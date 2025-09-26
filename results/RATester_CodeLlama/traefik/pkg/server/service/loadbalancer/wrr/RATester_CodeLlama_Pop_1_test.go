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
			{name: "a"},
			{name: "b"},
			{name: "c"},
		},
	}
	h := b.Pop()
	if h == nil {
		t.Fatal("expected handler")
	}
	if h.(*namedHandler).name != "a" {
		t.Fatal("expected handler a")
	}
	if len(b.handlers) != 2 {
		t.Fatal("expected 2 handlers")
	}
	if b.handlers[0].name != "b" {
		t.Fatal("expected handler b")
	}
	if b.handlers[1].name != "c" {
		t.Fatal("expected handler c")
	}
}
