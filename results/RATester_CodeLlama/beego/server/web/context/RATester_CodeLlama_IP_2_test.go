package context

import (
	"fmt"
	"net/http"
	"testing"

	"gotest.tools/assert"
)

func TestIP_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				RemoteAddr: "127.0.0.1:8080",
			},
		},
	}
	assert.Equal(t, "127.0.0.1", input.IP())
}
