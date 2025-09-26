package testhelpers

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestWithSticky_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cookieName := "test"
	expected := &dynamic.ServersLoadBalancer{
		Sticky: &dynamic.Sticky{
			Cookie: &dynamic.Cookie{Name: cookieName},
		},
	}
	actual := &dynamic.ServersLoadBalancer{}
	WithSticky(cookieName)(actual)
	assert.Equal(t, expected, actual)
}
