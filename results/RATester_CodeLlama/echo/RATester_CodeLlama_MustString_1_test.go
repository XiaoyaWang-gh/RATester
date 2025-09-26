package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestMustString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	// GIVEN
	sourceParam := "sourceParam"
	dest := new(string)
	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			return "value"
		},
	}

	// WHEN
	b.MustString(sourceParam, dest)

	// THEN
	assert.Equal(t, "value", *dest)
}
