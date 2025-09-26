package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestApplyMiddleware_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	h := func(c Context) error {
		return nil
	}
	middleware := []MiddlewareFunc{
		func(next HandlerFunc) HandlerFunc {
			return func(c Context) error {
				return next(c)
			}
		},
		func(next HandlerFunc) HandlerFunc {
			return func(c Context) error {
				return next(c)
			}
		},
	}

	// when
	result := applyMiddleware(h, middleware...)

	// then
	assert.NotNil(t, result)
}
