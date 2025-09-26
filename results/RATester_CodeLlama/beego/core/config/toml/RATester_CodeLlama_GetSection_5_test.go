package toml

import (
	"fmt"
	"testing"

	"github.com/pelletier/go-toml"
)

func TestGetSection_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &configContainer{
		t: &toml.Tree{},
	}
	section := "a.b.c"
	res, err := c.GetSection(section)
	if err != nil {
		t.Errorf("GetSection() error = %v", err)
		return
	}
	if len(res) != 0 {
		t.Errorf("GetSection() res = %v, want %v", res, 0)
	}
}
