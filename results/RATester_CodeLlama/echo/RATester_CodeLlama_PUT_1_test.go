package echo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zeebo/assert"
)

func TestPUT_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	g := e.Group("/")
	g.PUT("/", func(c Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello, World!", rec.Body.String())
}
