package gin

import (
	"fmt"
	"testing"
)

func TestMatch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	group := &RouterGroup{}
	group.Match([]string{"GET", "POST"}, "/", func(c *Context) {})
	if len(group.Handlers) != 2 {
		t.Errorf("group.Handlers should have 2 elements")
	}
}
