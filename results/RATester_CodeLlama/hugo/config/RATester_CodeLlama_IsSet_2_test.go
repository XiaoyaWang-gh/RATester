package config

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestIsSet_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &defaultConfigProvider{}
	c.Set("key", "value")
	assert.True(t, c.IsSet("key"))
	assert.False(t, c.IsSet("key1"))
}
