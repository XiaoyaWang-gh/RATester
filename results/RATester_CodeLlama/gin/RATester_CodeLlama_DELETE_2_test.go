package gin

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-playground/assert"
)

func TestDELETE_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	group := &RouterGroup{}
	group.DELETE("/", func(c *Context) {
		c.String(http.StatusOK, "Hello, World!")
	})
	assert.Equal(t, 1, len(group.Handlers))
}
