package fiber

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestStack_1(t *testing.T) {
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
				Name: "route1",
			},
			{
				Name: "route2",
			},
		},
		{
			{
				Name: "route3",
			},
			{
				Name: "route4",
			},
		},
	}

	// Act
	stack := app.Stack()

	// Assert
	assert.Equal(t, stack, app.stack)
}
