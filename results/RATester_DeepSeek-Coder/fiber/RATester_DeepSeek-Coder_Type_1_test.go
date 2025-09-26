package fiber

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()

	app.Get("/", func(c Ctx) error {
		c.Type("json")
		return c.SendString("Hello, World 👋")
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Errorf("Test failed with error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %v", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json; charset=utf-8" {
		t.Errorf("Expected Content-Type 'application/json; charset=utf-8', got %v", contentType)
	}
}
