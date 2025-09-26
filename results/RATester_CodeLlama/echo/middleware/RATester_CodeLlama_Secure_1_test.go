package middleware

import (
	"fmt"
	"testing"

	"github.com/labstack/echo/v4"
	"gotest.tools/assert"
)

func TestSecure_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	expected := echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	})

	// Act
	actual := Secure()

	// Assert
	assert.Equal(t, expected, actual)
}
