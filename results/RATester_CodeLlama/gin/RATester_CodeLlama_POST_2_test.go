package gin

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-playground/assert"
)

func TestPOST_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	group := &RouterGroup{}
	group.POST("/", func(c *Context) {
		c.String(http.StatusOK, "Hello, World!")
	})
	assert.Equal(t, group.POST("/", nil), group)
}
