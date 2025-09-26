package web

import (
	"fmt"
	"testing"
)

func TestNamespace_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n1 := &Namespace{
		handlers: &ControllerRegister{
			routers: make(map[string]*Tree),
		},
	}
	n2 := &Namespace{
		handlers: &ControllerRegister{
			routers: make(map[string]*Tree),
		},
	}
	n1.handlers.routers["GET"] = NewTree()
	n2.handlers.routers["GET"] = NewTree()

	n1.Namespace(n2)

	if len(n1.handlers.routers) != 1 {
		t.Errorf("Expected 1 router, got %d", len(n1.handlers.routers))
	}
	if _, ok := n1.handlers.routers["GET"]; !ok {
		t.Errorf("Expected router for GET method, but not found")
	}
}
