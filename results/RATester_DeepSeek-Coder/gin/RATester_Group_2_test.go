package gin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func TestGroup_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	group := e.Group("/group")
	h := func(c *Context) {
		c.String(200, "OK")
	}
	group.GET("/get", h)
	req, _ := http.NewRequest("GET", "/group/get", nil)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}
