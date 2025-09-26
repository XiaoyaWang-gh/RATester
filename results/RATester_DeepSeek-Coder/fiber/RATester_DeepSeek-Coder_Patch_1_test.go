package fiber

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPatch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()
	registering := &Registering{app: app, path: "/test"}

	handler := func(c Ctx) error {
		return c.SendString("Hello, World 👋")
	}

	registering.Patch(handler)

	req := httptest.NewRequest(http.MethodPatch, "/test", nil)
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

	expectedBody := "Hello, World 👋"
	if string(body) != expectedBody {
		t.Errorf("Expected body %q, got %q", expectedBody, string(body))
	}
}
