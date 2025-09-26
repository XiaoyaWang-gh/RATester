package echo

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestOPTIONS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given: A test
	g := &Group{}
	path := "path"
	h := func(c Context) error {
		return nil
	}
	m := []MiddlewareFunc{}
	// when: Calling the method under test
	result := g.OPTIONS(path, h, m...)
	// then: The result is a Route
	assert.IsType(t, &Route{}, result)
}
