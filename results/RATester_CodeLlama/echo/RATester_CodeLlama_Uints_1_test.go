package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestUints_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// GIVEN
	sourceParam := "sourceParam"
	dest := &[]uint{}
	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			return "1,2,3"
		},
	}
	// WHEN
	b.Uints(sourceParam, dest)
	// THEN
	assert.Equal(t, []uint{1, 2, 3}, *dest)
}
