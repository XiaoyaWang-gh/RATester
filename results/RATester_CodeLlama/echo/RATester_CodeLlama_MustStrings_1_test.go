package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestMustStrings_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	sourceParam := "sourceParam"
	dest := &[]string{}
	b := &ValueBinder{}

	// when
	b.MustStrings(sourceParam, dest)

	// then
	assert.Equal(t, []string{}, *dest)
}
