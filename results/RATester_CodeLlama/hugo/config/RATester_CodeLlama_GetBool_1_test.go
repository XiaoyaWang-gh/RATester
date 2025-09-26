package config

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetBool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &defaultConfigProvider{}
	c.Set("key", "true")
	assert.True(t, c.GetBool("key"))
}
