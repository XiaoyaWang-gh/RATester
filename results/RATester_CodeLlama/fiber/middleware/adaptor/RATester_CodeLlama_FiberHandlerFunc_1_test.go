package adaptor

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/zeebo/assert"
)

func TestFiberHandlerFunc_1(t *testing.T) {
	t.Parallel()
	t.Run("should return a http.HandlerFunc", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		// given
		h := func(fiber.Ctx) error {
			return nil
		}
		// when
		handler := FiberHandlerFunc(h)
		// then
		assert.NotNil(t, handler)
	})
}
