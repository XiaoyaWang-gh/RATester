package earlydata

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3"
	"gotest.tools/assert"
)

func TestConfigDefault_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	config := []Config{
		{
			IsEarlyData: func(c fiber.Ctx) bool {
				return true
			},
		},
	}
	// Act
	result := configDefault(config...)
	// Assert
	assert.Equal(t, result.IsEarlyData, config[0].IsEarlyData)
}
