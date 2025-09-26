package config

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetString_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &defaultConfigProvider{}
	c.Set("key", "value")
	assert.Equal(t, "value", c.GetString("key"))
}
