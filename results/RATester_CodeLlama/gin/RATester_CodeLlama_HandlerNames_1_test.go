package gin

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestHandlerNames_1(t *testing.T) {
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
	assert.Equal(t, []string{"func(c *Context)", "func(c *Context)", "func(c *Context)"}, c.HandlerNames())
}
