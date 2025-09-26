package gin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoRoute_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	assert.Empty(t, e.noRoute)
	e.NoRoute(func(c *Context) {
		c.String(http.StatusOK, "test")
	})
	assert.NotEmpty(t, e.noRoute)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	e.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "test", w.Body.String())
}
