package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestMustInt8_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// GIVEN
	sourceParam := "sourceParam"
	dest := new(int8)
	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			return "1"
		},
	}
	// WHEN
	b.MustInt8(sourceParam, dest)
	// THEN
	assert.Equal(t, int8(1), *dest)
}
