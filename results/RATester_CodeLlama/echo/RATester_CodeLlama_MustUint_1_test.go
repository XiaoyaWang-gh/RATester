package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestMustUint_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	// GIVEN
	sourceParam := "sourceParam"
	dest := uint(0)
	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			return "1"
		},
	}

	// WHEN
	b.MustUint(sourceParam, &dest)

	// THEN
	assert.Equal(t, uint(1), dest)
}
