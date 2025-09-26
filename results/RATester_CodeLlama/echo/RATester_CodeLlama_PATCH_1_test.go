package echo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zeebo/assert"
)

func TestPATCH_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	assert.NotNil(t, e)

	e.PATCH("/", func(c Context) error {
		return c.String(http.StatusOK, "test")
	})

	req := httptest.NewRequest(http.MethodPatch, "/", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "test", rec.Body.String())
}
