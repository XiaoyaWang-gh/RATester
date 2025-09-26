package gin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func TestOPTIONS_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	router := New()

	router.OPTIONS("/options", func(c *Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", "/options", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello, World!", w.Body.String())
}
