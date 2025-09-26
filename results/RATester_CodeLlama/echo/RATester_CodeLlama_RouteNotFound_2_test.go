package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestRouteNotFound_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	r := e.RouteNotFound("/", func(c Context) error {
		return nil
	})
	assert.NotNil(t, r)
}
