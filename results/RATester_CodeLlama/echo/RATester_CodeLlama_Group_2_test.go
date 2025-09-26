package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestGroup_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	g := e.Group("/")
	assert.NotNil(t, g)
	assert.Equal(t, "/", g.prefix)
	assert.Equal(t, e, g.echo)
}
