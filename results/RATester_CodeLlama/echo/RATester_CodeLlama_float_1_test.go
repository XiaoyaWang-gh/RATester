package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestFloat_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	// GIVEN
	sourceParam := "sourceParam"
	value := "1.23"
	dest := float64(0)
	bitSize := 64
	b := &ValueBinder{}

	// WHEN
	b.float(sourceParam, value, dest, bitSize)

	// THEN
	assert.Equal(t, float64(1.23), dest)
}
