package config

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetInt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &defaultConfigProvider{}
	c.Set("key", "1")
	assert.Equal(t, 1, c.GetInt("key"))
}
