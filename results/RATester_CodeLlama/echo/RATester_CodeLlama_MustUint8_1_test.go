package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestMustUint8_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	// GIVEN
	sourceParam := "sourceParam"
	dest := new(uint8)
	b := new(ValueBinder)
	b.ValueFunc = func(sourceParam string) string {
		return "1"
	}

	// WHEN
	b.MustUint8(sourceParam, dest)

	// THEN
	assert.Equal(t, uint8(1), *dest)
}
