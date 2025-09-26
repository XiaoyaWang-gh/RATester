package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGET_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	group := &RouterGroup{}
	group.GET("/", func(c *Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	if group.GET("/") == nil {
		t.Errorf("GET() = %v, want %v", group.GET("/"), nil)
	}
}
