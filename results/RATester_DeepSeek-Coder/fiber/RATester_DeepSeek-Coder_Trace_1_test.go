package fiber

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTrace_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()
	registering := &Registering{app: app, path: "/test"}

	handler := func(ctx Ctx) error {
		return ctx.SendString("Hello, World")
	}

	registering.Trace(handler)

	req := httptest.NewRequest("TRACE", "/test", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
