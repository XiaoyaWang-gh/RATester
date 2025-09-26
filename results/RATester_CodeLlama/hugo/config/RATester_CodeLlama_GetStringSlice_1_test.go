package config

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetStringSlice_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &defaultConfigProvider{}
	c.Set("key", []string{"a", "b", "c"})
	assert.Equal(t, []string{"a", "b", "c"}, c.GetStringSlice("key"))
}
