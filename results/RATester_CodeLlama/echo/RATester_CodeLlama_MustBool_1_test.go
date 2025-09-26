package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestMustBool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	sourceParam := "sourceParam"
	dest := new(bool)
	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			return "true"
		},
	}

	// when
	b.MustBool(sourceParam, dest)

	// then
	assert.Equal(t, true, *dest)
}
