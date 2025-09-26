package toml

import (
	"fmt"
	"testing"

	"github.com/pelletier/go-toml"
)

func TestDefaultInt64_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &configContainer{
		t: &toml.Tree{},
	}
	key := "a.b.c"
	defaultVal := int64(100)
	res := c.DefaultInt64(key, defaultVal)
	if res != defaultVal {
		t.Errorf("DefaultInt64 failed")
	}
}
