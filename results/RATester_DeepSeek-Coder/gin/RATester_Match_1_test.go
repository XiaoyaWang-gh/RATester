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

	router.Match([]string{"GET", "POST"}, "/path", func(c *Context) {
		c.String(200, "Hello, %s", "World")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/path", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Hello, World", w.Body.String())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/path", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Hello, World", w.Body.String())
}
