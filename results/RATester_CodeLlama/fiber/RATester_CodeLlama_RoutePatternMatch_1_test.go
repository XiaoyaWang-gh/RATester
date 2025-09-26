package fiber

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestRoutePatternMatch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	path := "path"
	pattern := "pattern"
	cfg := []Config{}

	// Act
	actual := RoutePatternMatch(path, pattern, cfg...)

	// Assert
	assert.Equal(t, true, actual)
}
