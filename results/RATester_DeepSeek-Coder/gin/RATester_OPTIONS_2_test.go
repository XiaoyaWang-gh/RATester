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
		c.String(200, "OK")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", "/options", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}
