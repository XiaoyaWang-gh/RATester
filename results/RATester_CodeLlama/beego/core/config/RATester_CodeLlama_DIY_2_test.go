package config

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestDIY_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &IniConfigContainer{}
	c.data = make(map[string]map[string]string)
	c.data["section"] = make(map[string]string)
	c.data["section"]["key"] = "value"
	v, err := c.DIY("section:key")
	assert.NoError(t, err)
	assert.Equal(t, "value", v)
}
