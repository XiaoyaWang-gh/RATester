package config

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/common/maps"
)

func TestGet_18(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &defaultConfigProvider{
		root: maps.Params{
			"key1": "value1",
			"key2": "value2",
		},
	}
	assert.Equal(t, "value1", c.Get("key1"))
	assert.Equal(t, "value2", c.Get("key2"))
	assert.Nil(t, c.Get("key3"))
}
