package echo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHEAD_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()

	e.HEAD("/test", func(c Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	req := httptest.NewRequest(http.MethodHead, "/test", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	if rec.Body.String() != "" {
		t.Errorf("Expected empty body, got %s", rec.Body.String())
	}
}
