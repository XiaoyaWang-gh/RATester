package middleware

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
)

func TestContextTimeoutWithConfig_1(t *testing.T) {
	t.Run("should return middleware function when config is valid", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		config := ContextTimeoutConfig{
			Skipper: func(c echo.Context) bool {
				return false
			},
			ErrorHandler: func(err error, c echo.Context) error {
				return c.String(http.StatusRequestTimeout, "Request Timeout")
			},
			Timeout: 1 * time.Second,
		}
		mw, err := config.ToMiddleware()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if mw == nil {
			t.Errorf("Expected middleware function, got nil")
		}
	})

	t.Run("should panic when config is invalid", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		config := ContextTimeoutConfig{
			Skipper: func(c echo.Context) bool {
				return false
			},
			ErrorHandler: func(err error, c echo.Context) error {
				return c.String(http.StatusRequestTimeout, "Request Timeout")
			},
			Timeout: -1 * time.Second,
		}
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic, got none")
			}
		}()
		_, err := config.ToMiddleware()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
