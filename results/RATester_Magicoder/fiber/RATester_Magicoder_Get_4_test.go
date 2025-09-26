package fiber

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()
	registering := &Registering{app: app, path: "/test"}

	handler := func(ctx Ctx) error {
		return ctx.SendString("Hello, World!")
	}

	registering.Get(handler)

	req := httptest.NewRequest("GET", "/test", nil)
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

	if string(body) != "Hello, World!" {
		t.Errorf("Expected body 'Hello, World!', got '%s'", string(body))
	}
}
