package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestIntValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	// GIVEN
	sourceParam := "sourceParam"
	dest := 1
	bitSize := 32
	valueMustExist := true
	b := &ValueBinder{}

	// WHEN
	b.intValue(sourceParam, dest, bitSize, valueMustExist)

	// THEN
	assert.Equal(t, 1, dest)
}
