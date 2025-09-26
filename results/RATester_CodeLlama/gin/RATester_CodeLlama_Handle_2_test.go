package gin

import (
	"fmt"
	"testing"
)

func TestHandle_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	group := &RouterGroup{}
	group.Handle("GET", "/", func(c *Context) {})
	if len(group.Handlers) != 1 {
		t.Errorf("group.Handle() failed")
	}
}
