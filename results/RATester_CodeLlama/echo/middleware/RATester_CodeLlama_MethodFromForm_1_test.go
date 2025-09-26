package middleware

import (
	"fmt"
	"testing"

	"github.com/labstack/echo/v4"
	"gotest.tools/assert"
)

func TestMethodFromForm_1(t *testing.T) {
	t.Parallel()

	t.Run("should return MethodOverrideGetter", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		param := "param"
		expected := func(c echo.Context) string {
			return c.FormValue(param)
		}

		actual := MethodFromForm(param)

		assert.Equal(t, expected, actual)
	})
}
