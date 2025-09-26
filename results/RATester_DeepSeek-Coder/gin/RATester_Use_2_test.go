package gin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUse_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := New()

	// Define middleware
	middleware1 := func(c *Context) {
		c.Set("middleware1", "true")
		c.Next()
	}

	middleware2 := func(c *Context) {
		c.Set("middleware2", "true")
		c.Next()
	}

	// Use middleware
	engine.Use(middleware1, middleware2)

	// Define route
	engine.GET("/test", func(c *Context) {
		// Check if middleware was applied
		m1, exists1 := c.Get("middleware1")
		m2, exists2 := c.Get("middleware2")

		if !exists1 || m1 != "true" || !exists2 || m2 != "true" {
			t.Errorf("Expected middleware to be applied")
		}
	})

	// Perform request
	req, _ := http.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	// Check if middleware was applied
	if w.Body.String() != "true" {
		t.Errorf("Expected middleware to be applied")
	}
}
