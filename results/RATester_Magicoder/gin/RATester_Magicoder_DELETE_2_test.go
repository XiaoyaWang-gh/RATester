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

	router.DELETE("/user/:name", func(c *Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/user/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello test", w.Body.String())
}
