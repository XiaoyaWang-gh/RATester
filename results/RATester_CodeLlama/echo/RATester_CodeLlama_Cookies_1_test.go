package echo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zeebo/assert"
)

func TestCookies_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("foo", "bar")
	c.Set("baz", "bat")
	assert.Equal(t, []*http.Cookie{
		{
			Name:  "foo",
			Value: "bar",
		},
		{
			Name:  "baz",
			Value: "bat",
		},
	}, c.Cookies())
}
