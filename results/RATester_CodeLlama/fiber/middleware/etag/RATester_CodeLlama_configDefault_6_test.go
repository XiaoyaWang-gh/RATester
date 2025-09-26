package etag

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3"
	"gotest.tools/assert"
)

func TestConfigDefault_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	config := []Config{
		{
			Next: func(c fiber.Ctx) bool {
				return true
			},
			Weak: true,
		},
	}

	// Act
	actual := configDefault(config...)

	// Assert
	assert.Equal(t, config[0], actual)
}
