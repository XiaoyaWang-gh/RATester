package echo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zeebo/assert"
)

func TestTRACE_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	g := e.Group("/")
	g.TRACE("/", func(c Context) error {
		return c.String(http.StatusOK, "TRACE")
	})
	req := httptest.NewRequest(http.MethodTrace, "/", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "TRACE", rec.Body.String())
}
