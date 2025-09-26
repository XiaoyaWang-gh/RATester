package yaml

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestDefaultString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigContainer{
		data: map[string]interface{}{
			"key": "value",
		},
	}
	key := "key"
	defaultVal := "default"
	assert.Equal(t, "value", c.DefaultString(key, defaultVal))
	key = "key1"
	assert.Equal(t, "default", c.DefaultString(key, defaultVal))
}
