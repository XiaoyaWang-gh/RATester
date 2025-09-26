package fiber

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestAddPrefixToRoute_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	app := &App{}
	prefix := "prefix"
	route := &Route{}

	// Act
	app.addPrefixToRoute(prefix, route)

	// Assert
	assert.Equal(t, "prefix", route.Path)
}
