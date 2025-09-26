package echo

import (
	"fmt"
	"testing"
	"time"

	"github.com/zeebo/assert"
)

func TestMustUnixTimeMilli_1(t *testing.T) {
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
			return "1643298240000"
		},
	}
	// WHEN
	b.MustUnixTimeMilli(sourceParam, dest)
	// THEN
	assert.Equal(t, time.Unix(1643298240, 0), *dest)
}
