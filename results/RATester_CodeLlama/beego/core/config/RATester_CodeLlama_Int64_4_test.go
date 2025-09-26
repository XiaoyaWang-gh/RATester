package config

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestInt64_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &IniConfigContainer{}
	c.data = map[string]map[string]string{
		"section": {
			"key": "123",
		},
	}
	v, err := c.Int64("section:key")
	assert.NoError(t, err)
	assert.Equal(t, int64(123), v)
}
