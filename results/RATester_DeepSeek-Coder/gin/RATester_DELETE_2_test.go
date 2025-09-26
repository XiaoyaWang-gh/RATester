package gin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func TestDELETE_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	router := New()

	router.DELETE("/user/:id", func(c *Context) {
		c.String(200, "User %s deleted", c.Param("id"))
	})

	req, _ := http.NewRequest("DELETE", "/user/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Equal(t, "User 1 deleted", resp.Body.String())
}
