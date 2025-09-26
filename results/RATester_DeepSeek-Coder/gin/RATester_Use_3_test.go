package gin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUse_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	router := New()

	router.Use(func(c *Context) {
		c.Set("key", "value")
		c.Next()
	})

	router.GET("/", func(c *Context) {
		value, exists := c.Get("key")
		if !exists {
			t.Errorf("Expected key to exist in context")
		}
		if value != "value" {
			t.Errorf("Expected value to be 'value', got %v", value)
		}
	})

	router.GET("/", func(c *Context) {
		value, exists := c.Get("key")
		if !exists {
			t.Errorf("Expected key to exist in context")
		}
		if value != "value" {
			t.Errorf("Expected value to be 'value', got %v", value)
		}
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)
}
