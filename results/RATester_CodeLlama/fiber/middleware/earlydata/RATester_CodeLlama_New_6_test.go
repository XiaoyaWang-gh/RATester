package earlydata

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/zeebo/assert"
)

func TestNew_6(t *testing.T) {
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
			IsEarlyData: func(c fiber.Ctx) bool {
				return true
			},
			AllowEarlyData: func(c fiber.Ctx) bool {
				return true
			},
			Error: fiber.ErrTooEarly,
		},
	}

	// Act
	handler := New(config...)

	// Assert
	assert.NotNil(t, handler)
}
