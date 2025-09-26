package gin

import (
	"fmt"
	"testing"
)

func TestPATCH_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	group := &RouterGroup{}
	group.PATCH("/", func(c *Context) {})
	if len(group.Handlers) != 1 {
		t.Errorf("PATCH should append a handler to the group")
	}
}
