package fiber

import (
	"fmt"
	"testing"
)

func TestSetIndexRoute_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var c DefaultCtx
	var route int
	// Act
	c.setIndexRoute(route)
	// Assert
	if c.indexRoute != route {
		t.Errorf("Expected %v, got %v", route, c.indexRoute)
	}
}
