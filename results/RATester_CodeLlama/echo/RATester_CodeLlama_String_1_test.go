package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	sourceParam := "sourceParam"
	dest := ""
	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			return "value"
		},
	}

	// when
	b.String(sourceParam, &dest)

	// then
	assert.Equal(t, "value", dest)
}
