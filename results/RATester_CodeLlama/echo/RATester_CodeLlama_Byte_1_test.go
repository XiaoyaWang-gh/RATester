package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestByte_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	// GIVEN
	sourceParam := "sourceParam"
	dest := byte(0)
	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			return "1"
		},
	}

	// WHEN
	b.Byte(sourceParam, &dest)

	// THEN
	assert.Equal(t, byte(1), dest)
}
