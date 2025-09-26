package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestStrings_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	sourceParam := "sourceParam"
	dest := &[]string{}
	b := &ValueBinder{
		ValuesFunc: func(sourceParam string) []string {
			return []string{"1", "2"}
		},
	}

	// when
	b.Strings(sourceParam, dest)

	// then
	assert.Equal(t, []string{"1", "2"}, *dest)
}
