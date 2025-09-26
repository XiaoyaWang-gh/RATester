package toml

import (
	"fmt"
	"testing"

	"github.com/pelletier/go-toml"
	"github.com/zeebo/assert"
)

func TestString_26(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &configContainer{
		t: &toml.Tree{},
	}
	c.t.Set("a.b.c", "123")
	res, err := c.String("a.b.c")
	assert.NoError(t, err)
	assert.Equal(t, "123", res)
}
