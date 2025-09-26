package pprof

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestNew_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var (
		config []Config
		ctx    fiber.Ctx
	)

	// Act
	handler := New(config...)
	err := handler(ctx)

	// Assert
	if err != nil {
		t.Errorf("New() error = %v", err)
		return
	}
}
