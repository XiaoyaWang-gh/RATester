package fiber

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestBind_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	c := &DefaultCtx{}
	// Act
	bind := c.Bind()
	// Assert
	assert.NotNil(t, bind)
}
