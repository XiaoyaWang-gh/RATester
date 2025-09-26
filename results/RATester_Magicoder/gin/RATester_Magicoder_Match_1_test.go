package gin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func TestMatch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	router := New()
	group := router.Group("/group")

	group.Match([]string{"GET", "POST"}, "/path", func(c *Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/group/path", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello, World!", w.Body.String())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/group/path", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello, World!", w.Body.String())
}
