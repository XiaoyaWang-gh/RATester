package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestInSection_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var pt pageTree
	var other any

	// when
	result := pt.InSection(other)

	// then
	assert.False(t, result)
}
