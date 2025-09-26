package gin

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestUse_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	group := &RouterGroup{}
	group.Use(func(c *Context) {})
	assert.Equal(t, 1, len(group.Handlers))
}
