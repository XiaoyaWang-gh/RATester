package toml

import (
	"fmt"
	"testing"

	"github.com/pelletier/go-toml"
	"github.com/zeebo/assert"
)

func TestBool_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &configContainer{
		t: &toml.Tree{},
	}
	c.t.Set("a.b.c", true)
	b, err := c.Bool("a.b.c")
	assert.NoError(t, err)
	assert.Equal(t, true, b)
}
