package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/core/config"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshaler_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: setup test
	b := &beegoAppConfig{}
	prefix := "TODO: provide value for parameter `prefix`"
	obj := "TODO: provide value for parameter `obj`"
	opt := []config.DecodeOption{config.DecodeOption(nil)}
	err := b.Unmarshaler(prefix, obj, opt...)
	assert.NoError(t, err)
	// TODO: assert results
}
