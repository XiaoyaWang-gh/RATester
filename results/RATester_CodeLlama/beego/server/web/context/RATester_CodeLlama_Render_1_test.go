package context

import (
	"fmt"
	"testing"
)

func TestRender_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var f rendererFunc
	var ctx *Context
	f.Render(ctx)
}
