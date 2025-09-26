package toml

import (
	"fmt"
	"testing"

	"github.com/pelletier/go-toml"
)

func TestDefaultFloat_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &configContainer{
		t: &toml.Tree{},
	}
	key := "a.b.c"
	defaultVal := float64(1.1)
	res := c.DefaultFloat(key, defaultVal)
	if res != defaultVal {
		t.Errorf("DefaultFloat() = %v, want %v", res, defaultVal)
	}
}
