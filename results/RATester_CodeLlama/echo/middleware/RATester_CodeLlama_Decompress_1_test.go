package middleware

import (
	"fmt"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestDecompress_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var (
		nextHandlerFunc = func(c echo.Context) error {
			return nil
		}
		middlewareFunc = Decompress()
	)

	// Act
	middlewareFunc(nextHandlerFunc)(nil)

	// Assert
	// TODO
}
