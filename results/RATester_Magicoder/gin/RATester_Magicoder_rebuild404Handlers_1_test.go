package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func Testrebuild404Handlers_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := &Engine{
		noRoute: HandlersChain{
			func(c *Context) {
				c.String(http.StatusOK, "Hello, World!")
			},
		},
	}

	engine.rebuild404Handlers()

	if len(engine.allNoRoute) != 1 {
		t.Errorf("Expected allNoRoute to have 1 handler, but got %d", len(engine.allNoRoute))
	}

	handler := engine.allNoRoute[0]
	handler(nil)
}
