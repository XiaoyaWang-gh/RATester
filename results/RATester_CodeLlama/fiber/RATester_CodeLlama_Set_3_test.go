package fiber

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestSet_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	c := &DefaultCtx{}
	key := "key"
	val := "val"

	// Act
	c.Set(key, val)

	// Assert
	assert.Equal(t, val, c.fasthttp.Response.Header.Peek(key))
}
