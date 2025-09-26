package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHandlerNames_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		handlers: HandlersChain{
			func(c *Context) {
				// some logic
			},
			func(c *Context) {
				// some logic
			},
		},
	}

	expected := []string{"func1", "func2"}
	result := ctx.HandlerNames()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
