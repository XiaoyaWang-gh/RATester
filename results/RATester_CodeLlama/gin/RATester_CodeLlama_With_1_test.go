package gin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWith_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := New()
	engine.With(func(engine *Engine) {
		engine.Use(func(c *Context) {
			c.String(http.StatusOK, "test")
		})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	engine.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	if w.Body.String() != "test" {
		t.Errorf("Expected body %s but got %s", "test", w.Body.String())
	}
}
