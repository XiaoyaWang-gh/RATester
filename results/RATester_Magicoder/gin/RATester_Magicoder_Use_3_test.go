package gin

import (
	"fmt"
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
			t.Error("Expected key to exist in context")
		}
		if value != "value" {
			t.Error("Expected value to be 'value'")
		}
	})

	router.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil))
}
