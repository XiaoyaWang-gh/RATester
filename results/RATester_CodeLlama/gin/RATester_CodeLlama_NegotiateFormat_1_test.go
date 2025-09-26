package gin

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-playground/assert"
)

func TestNegotiateFormat_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Request: &http.Request{
			Header: http.Header{
				"Accept": []string{"text/html"},
			},
		},
	}
	assert.Equal(t, "text/html", c.NegotiateFormat("text/html", "text/plain"))
}
