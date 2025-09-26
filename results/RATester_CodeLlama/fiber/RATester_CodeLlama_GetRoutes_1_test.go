package fiber

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestGetRoutes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	app := &App{}
	app.stack = [][]*Route{
		{
			{
				path: "/",
			},
			{
				path: "/test",
			},
		},
		{
			{
				path: "/test",
			},
		},
	}
	// Act
	routes := app.GetRoutes()
	// Assert
	assert.Equal(t, 2, len(routes))
	assert.Equal(t, "/", routes[0].path)
	assert.Equal(t, "/test", routes[1].path)
}
