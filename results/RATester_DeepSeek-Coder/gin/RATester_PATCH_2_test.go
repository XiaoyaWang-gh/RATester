package gin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func TestPATCH_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	router := New()

	router.PATCH("/patch", func(c *Context) {
		c.String(200, "PATCH method on /patch")
	})

	req, _ := http.NewRequest("PATCH", "/patch", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "PATCH method on /patch", w.Body.String())
}
