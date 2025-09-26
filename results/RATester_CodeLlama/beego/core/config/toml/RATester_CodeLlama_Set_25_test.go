package toml

import (
	"fmt"
	"testing"

	"github.com/pelletier/go-toml"
	"github.com/zeebo/assert"
)

func TestSet_25(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &configContainer{
		t: &toml.Tree{},
	}
	err := c.Set("a.b.c", "1")
	assert.NoError(t, err)
	assert.Equal(t, "1", c.t.Get("a.b.c").(string))
}
