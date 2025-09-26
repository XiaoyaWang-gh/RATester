package echo

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/zeebo/assert"
)

func TestQueryParam_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &context{
		request: &http.Request{
			URL: &url.URL{
				RawQuery: "name=john",
			},
		},
	}
	assert.Equal(t, "john", c.QueryParam("name"))
}
