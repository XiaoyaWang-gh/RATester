package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func Testrebuild405Handlers_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := &Engine{
		HandleMethodNotAllowed: true,
		noMethod: []HandlerFunc{
			func(c *Context) {
				c.String(http.StatusOK, "Hello, World!")
			},
		},
	}

	engine.rebuild405Handlers()

	if len(engine.allNoMethod) != 1 {
		t.Errorf("Expected 1 handler, got %d", len(engine.allNoMethod))
	}
}
