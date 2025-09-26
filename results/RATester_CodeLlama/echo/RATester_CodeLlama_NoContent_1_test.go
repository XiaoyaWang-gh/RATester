package echo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zeebo/assert"
)

func TestNoContent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, c.NoContent(http.StatusOK))
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 0, len(rec.Body.Bytes()))
}
