package recover //nolint:predeclared // TODO: Rename to some non-builtin
import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/zeebo/assert"
)

func TestNew_23(t *testing.T) {
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
		},
	}
	// Act
	handler := New(config...)
	// Assert
	assert.NotNil(t, handler)
}
