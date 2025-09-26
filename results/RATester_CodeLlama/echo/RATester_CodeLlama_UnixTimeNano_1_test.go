package echo

import (
	"fmt"
	"testing"
	"time"

	"github.com/zeebo/assert"
)

func TestUnixTimeNano_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// GIVEN
	sourceParam := "sourceParam"
	dest := &time.Time{}
	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			return "1631020800000000000"
		},
	}
	// WHEN
	b.UnixTimeNano(sourceParam, dest)
	// THEN
	assert.Equal(t, time.Unix(1631020800, 0), *dest)
}
