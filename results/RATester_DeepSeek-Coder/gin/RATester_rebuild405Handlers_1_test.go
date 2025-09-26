package gin

import (
	"fmt"
	"testing"
)

func TestRebuild405Handlers_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := &Engine{
		noMethod: []HandlerFunc{
			func(c *Context) {
				c.String(200, "OK")
			},
		},
	}

	engine.rebuild405Handlers()

	if len(engine.allNoMethod) != 1 {
		t.Errorf("Expected 1 handler, got %d", len(engine.allNoMethod))
	}

	engine.noMethod = append(engine.noMethod, func(c *Context) {
		c.String(200, "OK")
	})

	engine.rebuild405Handlers()

	if len(engine.allNoMethod) != 2 {
		t.Errorf("Expected 2 handlers, got %d", len(engine.allNoMethod))
	}
}
