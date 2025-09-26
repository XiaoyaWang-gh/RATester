package config

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestBool_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &IniConfigContainer{}
	c.data = map[string]map[string]string{
		"section": {
			"key": "true",
		},
	}
	b, err := c.Bool("section:key")
	assert.NoError(t, err)
	assert.True(t, b)
}
