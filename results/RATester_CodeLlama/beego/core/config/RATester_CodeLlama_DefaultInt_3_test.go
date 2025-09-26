package config

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestDefaultInt_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &fakeConfigContainer{
		data: map[string]string{
			"key": "1",
		},
	}
	key := "key"
	defaultVal := 2
	assert.Equal(t, 1, c.DefaultInt(key, defaultVal))
}
