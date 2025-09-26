package echo

import (
	"fmt"
	"net/http"
	"testing"

	"gotest.tools/assert"
)

func TestAdd_2(t *testing.T) {
	e := New()
	g := e.Group("/group")
	h := func(c Context) error {
		return c.String(http.StatusOK, "OK")
	}

	t.Run("Add", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		r := g.Add(http.MethodGet, "/test", h)
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/group/test", r.Path)
	})

	t.Run("AddWithMiddleware", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		m := func(next HandlerFunc) HandlerFunc {
			return func(c Context) error {
				return next(c)
			}
		}
		r := g.Add(http.MethodGet, "/test", h, m)
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/group/test", r.Path)
	})
}
