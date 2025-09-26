package gin

import (
	"fmt"
	"testing"
)

func TestLast_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	chain := HandlersChain{
		func(c *Context) {},
		func(c *Context) {},
		func(c *Context) {},
	}

	last := chain.Last()

	if last == nil {
		t.Error("Expected a HandlerFunc, but got nil")
	}
}
