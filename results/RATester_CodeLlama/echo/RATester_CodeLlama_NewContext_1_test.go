package echo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	r, _ := http.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	assert.NotNil(t, c)
	assert.Equal(t, r, c.Request())
	assert.Equal(t, w, c.Response())
}
