package toml

import (
	"fmt"
	"testing"

	"github.com/pelletier/go-toml"
)

func TestFloat_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &configContainer{
		t: &toml.Tree{},
	}
	c.t.Set("a.b.c", 1.2)
	res, err := c.Float("a.b.c")
	if err != nil {
		t.Error(err)
	}
	if res != 1.2 {
		t.Errorf("expect 1.2, but got %v", res)
	}
}
