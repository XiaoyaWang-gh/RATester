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

	b := &Balancer{}
	b.handlers = []*namedHandler{
		{name: "a"},
		{name: "b"},
		{name: "c"},
	}
	b.Swap(0, 1)
	if b.handlers[0].name != "b" {
		t.Errorf("expected b, got %s", b.handlers[0].name)
	}
	if b.handlers[1].name != "a" {
		t.Errorf("expected a, got %s", b.handlers[1].name)
	}
	b.Swap(1, 2)
	if b.handlers[1].name != "c" {
		t.Errorf("expected c, got %s", b.handlers[1].name)
	}
	if b.handlers[2].name != "b" {
		t.Errorf("expected b, got %s", b.handlers[2].name)
	}
}
