package gin

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-playground/assert"
)

func TestRemoteIP_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Request: &http.Request{
			RemoteAddr: "192.168.1.1:8080",
		},
	}
	assert.Equal(t, "192.168.1.1", c.RemoteIP())
}
