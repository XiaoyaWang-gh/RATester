package config

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/common/maps"
)

func TestKeys_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &defaultConfigProvider{
		root: maps.Params{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	}
	assert.Equal(t, []string{"a", "b", "c"}, c.Keys())
}
