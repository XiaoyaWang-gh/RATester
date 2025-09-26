package gin

import (
	"fmt"
	"testing"
)

func TestAny_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	group := &RouterGroup{}
	group.Any("", func(c *Context) {})
	if len(group.Handlers) != 1 {
		t.Errorf("group.Handlers should have 1 element")
	}
}
