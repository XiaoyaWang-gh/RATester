package adaptor

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestHTTPHandlerFunc_1(t *testing.T) {
	t.Run("should return a fiber.Handler", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		handler := HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		if handler == nil {
			t.Errorf("Expected HTTPHandlerFunc to return a fiber.Handler, but got nil")
		}
	})

	t.Run("should correctly handle HTTP requests", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		app := fiber.New()
		app.Get("/", HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		resp, err := app.Test(httptest.NewRequest("GET", "/", nil))
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %v, but got %v", http.StatusOK, resp.StatusCode)
		}
	})
}
