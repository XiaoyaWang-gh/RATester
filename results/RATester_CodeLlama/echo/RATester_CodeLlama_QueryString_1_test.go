package echo

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/zeebo/assert"
)

func TestQueryString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &context{
		request: &http.Request{
			URL: &url.URL{
				RawQuery: "foo=bar",
			},
		},
	}
	assert.Equal(t, "foo=bar", c.QueryString())
}
