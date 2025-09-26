package config

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestString_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &IniConfigContainer{}
	c.Set("key", "val")
	val, err := c.String("key")
	assert.NoError(t, err)
	assert.Equal(t, "val", val)
}
