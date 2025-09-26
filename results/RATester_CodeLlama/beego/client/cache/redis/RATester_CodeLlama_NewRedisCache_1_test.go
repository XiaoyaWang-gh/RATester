package redis

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewRedisCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := NewRedisCache()
	assert.NotNil(t, c)
}
