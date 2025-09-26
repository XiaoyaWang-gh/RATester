package echo

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"

	"github.com/zeebo/assert"
)

func TestIsTLS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &context{
		request: &http.Request{
			TLS: &tls.ConnectionState{},
		},
	}

	// when
	result := c.IsTLS()

	// then
	assert.True(t, result)
}
