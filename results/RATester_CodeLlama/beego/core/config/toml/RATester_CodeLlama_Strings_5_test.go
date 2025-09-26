package toml

import (
	"fmt"
	"testing"

	"github.com/pelletier/go-toml"
	"github.com/zeebo/assert"
)

func TestStrings_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &configContainer{
		t: &toml.Tree{},
	}
	c.t.Set("a.b.c", []interface{}{"1", "2", "3"})
	val, err := c.Strings("a.b.c")
	assert.NoError(t, err)
	assert.Equal(t, []string{"1", "2", "3"}, val)
}
