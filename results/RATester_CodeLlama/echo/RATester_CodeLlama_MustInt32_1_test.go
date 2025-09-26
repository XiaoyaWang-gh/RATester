package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestMustInt32_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	// GIVEN
	sourceParam := "sourceParam"
	dest := new(int32)
	b := new(ValueBinder)
	b.ValueFunc = func(sourceParam string) string {
		return "1"
	}

	// WHEN
	b.MustInt32(sourceParam, dest)

	// THEN
	assert.Equal(t, int32(1), *dest)
}
