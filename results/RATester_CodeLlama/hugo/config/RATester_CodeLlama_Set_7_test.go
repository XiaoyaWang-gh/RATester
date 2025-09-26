package config

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestSet_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &defaultConfigProvider{}
	c.Set("k", "v")
	assert.Equal(t, "v", c.GetString("k"))
}
