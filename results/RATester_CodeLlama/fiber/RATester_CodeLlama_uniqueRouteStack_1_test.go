package fiber

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestUniqueRouteStack_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var stack []*Route
	// Act
	unique := uniqueRouteStack(stack)
	// Assert
	assert.Equal(t, []*Route{}, unique)
}
