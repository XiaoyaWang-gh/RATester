package echo

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/zeebo/assert"
)

func TestStatic_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	r := e.Static("/static", "./static")
	assert.NotNil(t, r)
	assert.Equal(t, http.MethodGet, r.Method)
	assert.Equal(t, "/static*", r.Path)
}
