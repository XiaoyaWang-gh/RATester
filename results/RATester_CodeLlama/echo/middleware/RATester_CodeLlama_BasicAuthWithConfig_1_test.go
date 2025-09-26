package middleware

import (
	"fmt"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/zeebo/assert"
)

func TestBasicAuthWithConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	config := BasicAuthConfig{
		Validator: func(username, password string, c echo.Context) (bool, error) {
			return true, nil
		},
	}

	// when
	middleware := BasicAuthWithConfig(config)

	// then
	assert.NotNil(t, middleware)
}
