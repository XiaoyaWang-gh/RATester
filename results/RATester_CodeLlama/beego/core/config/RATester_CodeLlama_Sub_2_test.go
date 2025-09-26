package config

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestSub_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &BaseConfiger{}
	_, err := c.Sub("a")
	assert.Equal(t, "unsupported operation", err.Error())
}
