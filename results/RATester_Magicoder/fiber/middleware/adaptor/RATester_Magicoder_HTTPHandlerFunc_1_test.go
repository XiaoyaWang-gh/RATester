package adaptor

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestHTTPHandlerFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}

	fiberHandler := HTTPHandlerFunc(http.HandlerFunc(handler))

	app := fiber.New()
	app.Get("/", fiberHandler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
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
		t.Errorf("Expected body %s, got %s", "Hello, World!", string(body))
	}
}
