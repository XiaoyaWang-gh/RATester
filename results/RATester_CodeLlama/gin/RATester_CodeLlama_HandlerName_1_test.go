package gin

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestHandlerName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		handlers: HandlersChain{
			func(c *Context) {},
			func(c *Context) {},
			func(c *Context) {},
		},
	}
	assert.Equal(t, "func(c *Context) {}", c.HandlerName())
}
