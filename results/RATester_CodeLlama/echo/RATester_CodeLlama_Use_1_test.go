package echo

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestUse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	g := e.Group("/")
	g.Use(func(next HandlerFunc) HandlerFunc {
		return func(c Context) error {
			return next(c)
		}
	})
	assert.Len(t, g.middleware, 1)
}
