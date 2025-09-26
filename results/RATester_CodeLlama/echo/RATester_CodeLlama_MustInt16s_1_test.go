package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestMustInt16s_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// GIVEN
	sourceParam := "sourceParam"
	dest := []int16{}
	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			return "1,2,3"
		},
	}
	// WHEN
	b.MustInt16s(sourceParam, &dest)
	// THEN
	assert.Equal(t, []int16{1, 2, 3}, dest)
}
