package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestMustInt32s_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// GIVEN
	sourceParam := "sourceParam"
	dest := &[]int32{}
	b := &ValueBinder{}
	// WHEN
	b.MustInt32s(sourceParam, dest)
	// THEN
	assert.Equal(t, []int32{}, *dest)
}
