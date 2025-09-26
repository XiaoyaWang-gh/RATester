package echo

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/zeebo/assert"
)

func TestFormValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &context{
		request: &http.Request{
			Form: url.Values{
				"name": []string{"value"},
			},
		},
	}
	assert.Equal(t, "value", c.FormValue("name"))
}
