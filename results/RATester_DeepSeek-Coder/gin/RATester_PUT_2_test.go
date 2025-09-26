package gin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func TestPUT_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	router := New()

	router.PUT("/path", func(c *Context) {
		c.String(200, "PUT method on /path")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/path", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "PUT method on /path", w.Body.String())
}
