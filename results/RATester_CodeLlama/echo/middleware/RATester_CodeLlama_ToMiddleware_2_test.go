package middleware

import (
	"fmt"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestToMiddleware_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	// GIVEN
	config := RequestLoggerConfig{
		LogValuesFunc: func(c echo.Context, v RequestLoggerValues) error {
			return nil
		},
	}

	// WHEN
	middleware, err := config.ToMiddleware()

	// THEN
	require.NoError(t, err)
	require.NotNil(t, middleware)
}
