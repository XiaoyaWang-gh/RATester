package render

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestPositionerSourceTarget_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	c := &hookBase{}
	c.getSourceSample = func() []byte {
		return []byte("test")
	}

	// Act
	result := c.PositionerSourceTarget()

	// Assert
	assert.Equal(t, []byte("test"), result)
}
