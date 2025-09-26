package logger

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3"
	"gotest.tools/assert"
)

func TestStatusColor_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	code := 200
	colors := fiber.Colors{
		Green: "green",
	}
	// Act
	actual := statusColor(code, colors)
	// Assert
	assert.Equal(t, "green", actual)
}
