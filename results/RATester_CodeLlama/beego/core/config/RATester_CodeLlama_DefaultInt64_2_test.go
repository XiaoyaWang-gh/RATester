package config

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestDefaultInt64_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &fakeConfigContainer{
		data: map[string]string{
			"key": "100",
		},
	}
	key := "key"
	defaultVal := int64(10)
	assert.Equal(t, int64(100), c.DefaultInt64(key, defaultVal))
	assert.Equal(t, int64(10), c.DefaultInt64("not_exist_key", defaultVal))
}
