package fiber

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zeebo/assert"
)

func TestUse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()
	group := app.Group("/group")

	group.Use("prefix", func(ctx Ctx) error {
		return ctx.Next()
	})

	group.Get("/test", func(ctx Ctx) error {
		return ctx.SendString("Hello, World!")
	})

	req := httptest.NewRequest("GET", "/group/test", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, resp.StatusCode, http.StatusOK)
}
