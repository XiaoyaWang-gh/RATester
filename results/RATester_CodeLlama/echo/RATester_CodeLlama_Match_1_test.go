package echo

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestMatch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	r := e.Match([]string{"GET", "POST"}, "/", func(c Context) error {
		return nil
	}, func(next HandlerFunc) HandlerFunc {
		return func(c Context) error {
			return nil
		}
	})
	assert.Len(t, r, 2)
}
