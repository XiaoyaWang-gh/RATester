package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestMustInt8s_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// GIVEN
	sourceParam := "sourceParam"
	dest := &[]int8{}
	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			return "1,2,3"
		},
	}
	// WHEN
	b.MustInt8s(sourceParam, dest)
	// THEN
	assert.Equal(t, []int8{1, 2, 3}, *dest)
}
