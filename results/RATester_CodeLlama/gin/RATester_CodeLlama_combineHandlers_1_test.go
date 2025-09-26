package gin

import (
	"fmt"
	"testing"
)

func TestCombineHandlers_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	group := &RouterGroup{
		Handlers: []HandlerFunc{
			func(c *Context) {},
			func(c *Context) {},
		},
	}
	handlers := []HandlerFunc{
		func(c *Context) {},
		func(c *Context) {},
	}
	mergedHandlers := group.combineHandlers(handlers)
	assert1(len(mergedHandlers) == 4, "too many handlers")
}
