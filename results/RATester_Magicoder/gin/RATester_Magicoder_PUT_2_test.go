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

	router.PUT("/user", func(c *Context) {
		c.String(http.StatusOK, "PUT method on /user")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/user", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "PUT method on /user", w.Body.String())
}
