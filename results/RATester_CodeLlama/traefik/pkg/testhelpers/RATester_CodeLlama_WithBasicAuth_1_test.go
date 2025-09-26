package testhelpers

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestWithBasicAuth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	auth := &dynamic.BasicAuth{
		Users: []string{"test:test"},
	}
	middleware := WithBasicAuth(auth)
	assert.NotNil(t, middleware)
	m := &dynamic.Middleware{}
	middleware(m)
	assert.Equal(t, auth, m.BasicAuth)
}
