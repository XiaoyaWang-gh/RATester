package echo

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/zeebo/assert"
)

func TestAdd_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	r := e.Add("GET", "/", func(c Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	assert.NotNil(t, r)
	assert.Equal(t, "GET", r.Method)
	assert.Equal(t, "/", r.Path)
}
