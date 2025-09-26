package fiber

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGroup_1(t *testing.T) {
	app := New()
	group := app.Group("/v1")

	t.Run("TestGroup", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		handler := func(ctx Ctx) error {
			return ctx.SendString("Hello, World v1")
		}

		group.Get("/hello", handler)

		req := httptest.NewRequest("GET", "/v1/hello", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		if string(body) != "Hello, World v1" {
			t.Errorf("Expected body 'Hello, World v1', got '%s'", string(body))
		}
	})
}
