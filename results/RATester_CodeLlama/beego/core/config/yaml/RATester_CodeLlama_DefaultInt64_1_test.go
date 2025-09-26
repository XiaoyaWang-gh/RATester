package yaml

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestDefaultInt64_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigContainer{
		data: map[string]interface{}{
			"key": int64(100),
		},
	}
	key := "key"
	defaultVal := int64(1000)
	assert.Equal(t, int64(100), c.DefaultInt64(key, defaultVal))
}
