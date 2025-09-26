package logger

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestMethodColor_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	method := "GET"
	colors := fiber.Colors{
		Cyan: "\u001b[96m",
	}
	// Act
	actual := methodColor(method, colors)
	// Assert
	expected := "\u001b[96m"
	if actual != expected {
		t.Errorf("methodColor(%s, %v) = %s; want %s", method, colors, actual, expected)
	}
}
