package gin

import (
	"fmt"
	"testing"
)

func TestRebuild404Handlers_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := &Engine{
		noRoute: HandlersChain{
			func(c *Context) {
				c.String(200, "Hello, World")
			},
		},
	}

	engine.rebuild404Handlers()

	if engine.allNoRoute == nil {
		t.Errorf("Expected allNoRoute to be not nil")
	}

	if len(engine.allNoRoute) != 1 {
		t.Errorf("Expected allNoRoute to have 1 handler, got %d", len(engine.allNoRoute))
	}

	engine.noRoute = nil
	engine.rebuild404Handlers()

	if engine.allNoRoute == nil {
		t.Errorf("Expected allNoRoute to be not nil")
	}

	if len(engine.allNoRoute) != 0 {
		t.Errorf("Expected allNoRoute to have 0 handlers, got %d", len(engine.allNoRoute))
	}
}
