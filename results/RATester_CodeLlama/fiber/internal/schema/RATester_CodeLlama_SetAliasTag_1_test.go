package schema

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestSetAliasTag_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	e := &Encoder{}
	tag := "alias"

	// Act
	e.SetAliasTag(tag)

	// Assert
	assert.Equal(t, tag, e.cache.tag)
}
