package gin

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-playground/assert"
)

func TestHEAD_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	group := &RouterGroup{}
	group.HEAD("/", func(c *Context) {
		c.String(http.StatusOK, "OK")
	})
	assert.Equal(t, 1, len(group.Handlers))
}
