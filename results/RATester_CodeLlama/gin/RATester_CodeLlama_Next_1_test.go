package gin

import (
	"fmt"
	"testing"
)

func TestNext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		handlers: []HandlerFunc{
			func(c *Context) {
				c.Next()
			},
			func(c *Context) {
				c.Next()
			},
			func(c *Context) {
				c.Next()
			},
		},
	}
	c.Next()
	if c.index != 3 {
		t.Errorf("Expected c.index to be 3, but got: %d", c.index)
	}
}
