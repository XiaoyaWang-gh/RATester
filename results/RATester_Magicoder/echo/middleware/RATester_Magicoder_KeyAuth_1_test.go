package middleware

import (
	"fmt"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestKeyAuth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fn := func(auth string, c echo.Context) (bool, error) {
		// Implement your validation logic here
		return true, nil
	}

	h := KeyAuth(fn)

	if h == nil {
		t.Error("Expected a non-nil handler, but got nil")
	}
}
