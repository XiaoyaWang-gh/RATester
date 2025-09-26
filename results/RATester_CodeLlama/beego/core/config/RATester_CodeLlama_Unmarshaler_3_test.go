package config

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestUnmarshaler_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	prefix := "prefix"
	obj := &struct{}{}
	opt := []DecodeOption{func(options decodeOptions) {}}

	// when
	err := Unmarshaler(prefix, obj, opt...)

	// then
	assert.NoError(t, err)
}
