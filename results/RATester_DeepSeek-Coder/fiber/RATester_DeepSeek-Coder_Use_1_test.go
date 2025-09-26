package fiber

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUse_1(t *testing.T) {
	t.Run("TestUse", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		app := New()
		group := &Group{
			app: app,
		}

		testHandler := func(ctx Ctx) error {
			return ctx.SendString("Hello, World")
		}

		group.Use("testPrefix", testHandler)

		req := httptest.NewRequest("GET", "/testPrefix", nil)
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

		if string(body) != "Hello, World" {
			t.Errorf("Expected body 'Hello, World', got '%s'", string(body))
		}
	})
}
