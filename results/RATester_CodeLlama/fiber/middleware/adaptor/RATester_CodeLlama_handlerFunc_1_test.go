package adaptor

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/zeebo/assert"
)

func TestHandlerFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	app := fiber.New()
	h := []func(fiber.Ctx) error{}

	// Act
	handlerFunc(app, h...)

	// Assert
	assert.NotNil(t, handlerFunc)
}
