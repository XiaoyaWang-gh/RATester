package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestURL_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	h := func(c Context) error {
		return nil
	}
	r := e.URL(h, "a", "b")
	assert.Equal(t, "/a/b", r)
}
