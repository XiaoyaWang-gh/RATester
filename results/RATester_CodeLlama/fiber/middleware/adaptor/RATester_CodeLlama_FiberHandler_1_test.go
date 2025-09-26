package adaptor

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestFiberHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	h := func(fiber.Ctx) error {
		return nil
	}
	// Act
	result := FiberHandler(h)
	// Assert
	if result == nil {
		t.Errorf("FiberHandler() = %v, want %v", result, nil)
	}
}
