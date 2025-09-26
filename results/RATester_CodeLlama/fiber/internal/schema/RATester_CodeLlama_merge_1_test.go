package schema

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestMerge_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	e := MultiError{}
	errors := MultiError{}
	// When
	e.merge(errors)
	// Then
	assert.Equal(t, MultiError{}, e)
}
