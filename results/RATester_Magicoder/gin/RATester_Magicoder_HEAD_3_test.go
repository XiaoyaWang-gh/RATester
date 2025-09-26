package gin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func TestHEAD_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	router := New()

	router.HEAD("/path", func(c *Context) {
		c.String(200, "Hello, World!")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("HEAD", "/path", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "", w.Body.String())
}
